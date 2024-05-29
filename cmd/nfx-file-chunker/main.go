package main

import (
	_ "github.com/u2takey/ffmpeg-go"
	"log"
	"net/http"
	"netflix.com/chunker/cmd/nfx-file-chunker/config"
	"netflix.com/chunker/internal/router"
	"netflix.com/chunker/internal/storage"
	_ "netflix.com/chunker/internal/storage"
)

func main() {
	cfg, err := config.ParseConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
		return
	}

	storage.ConfigureS3Client(cfg)

	r := router.NewRouter()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
