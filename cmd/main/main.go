package main

import (
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	err := ffmpeg.Input("./data/invs.mp4").
		Output("./data/out%d.mp4",
			ffmpeg.KwArgs{"f": "segment", "c": "copy", "segment_time": "10", "reset_timestamps": "1"}).
		OverWriteOutput().ErrorToStdOut().Run()

	if err != nil {
		fmt.Println(err)
	}
}
