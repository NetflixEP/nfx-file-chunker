package chunker

import (
	"fmt"
	"os"
	"path"
)

const HlsMasterFile = "master.m3u8"

func createHLSMasterFile(dataPath string) error {
	file, err := os.Create(path.Join(dataPath, HlsMasterFile))
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString("#EXTM3U\n")
	for _, item := range playlist {
		file.WriteString(fmt.Sprintf("#EXT-X-STREAM-INF:BANDWIDTH=%d,RESOLUTION=%s\n", item.bandwidth, item.rsn.getQuality()))
		file.WriteString(fmt.Sprintf("%d.m3u8\n", item.rsn.height))
	}

	return nil
}
