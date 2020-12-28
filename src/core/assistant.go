package core

import (
	"app/common"
	"path"
)

type pdfValue struct {
	file string
	zoom int
	start int
	end int
	compress int
}

type pdfAssistant struct {
	queue chan * pdfValue
}

var PdfAssistant *pdfAssistant

func init()  {
	PdfAssistant = &pdfAssistant{
		queue: make(chan *pdfValue, 1024),
	}
}

func (pdf *pdfAssistant) Do(file string, zoom, start, end, compress int) int {
	size := PdfSize(path.Join(common.Config.PDFFolderPath, file))
	if size <= 0 {
		return 0
	}

	pdf.queue <- &pdfValue{
		file: file,
		zoom: zoom,
		start: start,
		end: end,
		compress: compress,
	}

	return size
}

func (pdf *pdfAssistant) Run(ctx *common.ServerContext) {
	defer ctx.Done()
	ctx.Add()

	for {
		select {
		case pv := <-pdf.queue:
			common.Logger.Printf("got pdf %+v", pv)
			PdfParse(path.Join(common.Config.PDFFolderPath, pv.file), pv.zoom, pv.start, pv.end, pv.compress)
		case <-ctx.Quit():
			common.Logger.Print("pdf progress catch exit signal")
			return
		}
	}
}
