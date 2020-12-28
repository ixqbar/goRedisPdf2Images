package core

import (
	"app/common"
	"github.com/jonnywang/redcon2"
	"github.com/tidwall/redcon"
	"os"
	"path"
	"strconv"
	"syscall"
	"time"
)

var pid int
var runAtTime time.Time

func init() {
	pid = syscall.Getpid()
	runAtTime = time.Now()
}

func ExitServer() {
	p, err := os.FindProcess(pid)
	if err == nil {
		p.Signal(syscall.SIGTERM)
	}
}

func RunRedisServer(ctx *common.ServerContext) {
	defer ctx.Done()
	ctx.Add()

	rs := redcon2.NewRedconServeMux()
	rs.Handle("version", func(conn redcon.Conn, cmd redcon.Command) {
		conn.WriteBulkString(common.VERSION)
	})
	rs.Handle("pdf", func(conn redcon.Conn, cmd redcon.Command) {
		if len(cmd.Args) < 2 {
			conn.WriteError("err command args with pdf name")
			return
		}

		file := string(cmd.Args[1])

		s, err := os.Stat(path.Join(common.Config.PDFFolderPath, file))
		if err != nil || s.Size() == 0 {
			common.Logger.Print(err)
			conn.WriteError("err command args with pdf name")
			return
		}

		zoom := 2
		start := 0
		end := 0
		compress := 1

		if len(cmd.Args) >= 3 {
			n, err := strconv.Atoi(string(cmd.Args[2]))
			if err != nil {
				conn.WriteError("err command args with start")
				return
			}
			start = n
		}

		if len(cmd.Args) >= 4 {
			n, err := strconv.Atoi(string(cmd.Args[3]))
			if err != nil {
				conn.WriteError("err command args with end")
				return
			}
			end = n
		}

		if start < 0 || end < 0 || (end > 0 && start > end) {
			conn.WriteError("err command args with start or end")
			return
		}

		if len(cmd.Args) >= 5 {
			n, err := strconv.Atoi(string(cmd.Args[4]))
			if err != nil {
				conn.WriteError("err command args with end")
				return
			}
			zoom = n
		}

		if len(cmd.Args) >= 6 {
			n, err := strconv.Atoi(string(cmd.Args[5]))
			if err != nil {
				conn.WriteError("err command args with end")
				return
			}
			compress = n
		}

		size := PdfAssistant.Do(file, zoom, start, end, compress)

		conn.WriteInt(size)
	})
	rs.Handle("png", func(conn redcon.Conn, cmd redcon.Command) {
		if len(cmd.Args) != 2 {
			conn.WriteError("err command args with png name")
			return
		}

		file := string(cmd.Args[1])

		s, err := os.Stat(file)
		if err != nil || s.Size() == 0 {
			common.Logger.Print(err)
			conn.WriteError("err command args with png name")
			return
		}

		go PngCompress(file)

		conn.WriteInt(1)
	})

	go func() {
		common.Logger.Printf("run redis protocol server at %+v with pid=%d", common.Config.Address, pid)
		err := rs.Run(common.Config.Address)
		if err != nil {
			common.Logger.Print(err)
			rs = nil
			ExitServer()
		}
	}()

	select {
	case <-ctx.Quit():
		common.Logger.Print("redis server catch exit signal")
		if rs != nil {
			rs.Close()
		}
	}
}

func Run() error {
	ctx := common.NewServerContext()

	ctx.Set("startTime", runAtTime)

	go PdfAssistant.Run(ctx)
	go RunRedisServer(ctx)

	select {
	case <-ctx.Interrupt():
		common.Logger.Print("server interrupt")
		ctx.Cancel()
	}

	ctx.Wait()
	common.Logger.Printf("server uptime %v %v", runAtTime.Format("2006-01-02 15:04:05"), time.Now().Sub(runAtTime))
	common.Logger.Print("server exit")

	return nil
}
