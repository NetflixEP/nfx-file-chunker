package router

import (
	"io"
	"net/http"
	"netflix.com/chunker/internal/chunker"
	"netflix.com/chunker/internal/storage"
	"os"
	"path/filepath"
)

const TmpDataDirectory = "data"

func uploadContent(w http.ResponseWriter, r *http.Request) {
	//Вычитываем файл из запроса
	file, handler, err := r.FormFile("file")
	if err != nil {
		BadRequestResponse(w, r, err)
		return
	}
	defer file.Close()

	destinationPath := filepath.Join(TmpDataDirectory, handler.Filename)
	//Определяем путь, куда необходимо положить файл
	dst, err := os.Create(destinationPath)
	if err != nil {
		ServerErrorResponse(w, r, err)
		return
	}
	defer dst.Close()

	//Сохраняем файл по месту назначения
	_, err = io.Copy(dst, file)
	if err != nil {
		ServerErrorResponse(w, r, err)
		return
	}

	resultPath, err := chunker.TranscodeFile(destinationPath)
	if err != nil {
		ServerErrorResponse(w, r, err)
		return
	}

	bucketName, err := storage.UploadContentFiles(resultPath)
	if err != nil {
		ServerErrorResponse(w, r, err)
	}
	w.Write([]byte("File " + handler.Filename + " uploaded successfully\n Bucket: " + bucketName))
}
