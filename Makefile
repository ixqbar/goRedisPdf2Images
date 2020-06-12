TARGET=goRedisPdf2Image

all: mac

linux:
	GOOS=linux GOARCH=amd64 go build -a -o ../bin/${TARGET}_${@} ./main.go

mac:
	cd src && GOOS=darwin GOARCH=amd64 go build -a -o ../bin/${TARGET}_${@} ./main.go

clean:
	rm -rf ./bin/${TARGET}_*	
