package chunker

import "fmt"

var playlist = newHLSPlaylist()

type hlsPlaylist []content

type content struct {
	rsn       resolution
	bandwidth int
}

type resolution struct {
	width  int
	height int
}

func (r *resolution) getQuality() string {
	return fmt.Sprintf("%dx%d", r.width, r.height)
}

func newHLSPlaylist() hlsPlaylist {
	return hlsPlaylist{
		{rsn: resolution{width: 640, height: 360}, bandwidth: 375000},
		{rsn: resolution{width: 854, height: 480}, bandwidth: 750000},
		{rsn: resolution{width: 1280, height: 720}, bandwidth: 2000000},
	}
}
