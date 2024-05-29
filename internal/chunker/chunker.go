package chunker

import (
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"golang.org/x/sync/errgroup"
	"os"
	"path"
	"strings"
)

func TranscodeFile(filename string) (string, error) {
	var eg errgroup.Group

	contentDir := getDirectoryName(filename)
	resultPath := path.Join("data", contentDir)
	err := os.MkdirAll(resultPath, os.ModePerm)
	if err != nil {
		return "", err
	}
	for _, cnt := range playlist {
		cnt := cnt
		eg.Go(func() error {
			return ffmpeg.Input(filename).
				Output(path.Join(resultPath, fmt.Sprintf("%d.m3u8", cnt.rsn.height)),
					ffmpeg.KwArgs{
						"profile:v":     "baseline",
						"level":         "3.0",
						"s":             cnt.rsn.getQuality(),
						"start_number":  "0",
						"hls_time":      "10",
						"hls_list_size": "0",
						"f":             "hls",
					},
				).ErrorToStdOut().Run()
		})
	}

	if err := eg.Wait(); err != nil {
		return "", err
	}
	os.Remove(filename)
	if err := createHLSMasterFile(resultPath); err != nil {
		return "", err
	}
	return resultPath, nil
}

func getDirectoryName(filename string) string {
	base := path.Base(filename)
	name := strings.TrimSuffix(base, ".mp4")
	return name
}
