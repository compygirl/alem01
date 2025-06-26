package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	// "path/filepath"
	// "strings"
)

func main() {
	http.HandleFunc("/api/archive/information", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests are supported", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "No file part", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Check if the uploaded file is a ZIP archive
		fileType := handler.Header.Get("Content-Type")
		if fileType != "application/zip" {
			http.Error(w, "Uploaded file is not a valid ZIP archive", http.StatusBadRequest)
			return
		}

		archiveInfo := map[string]interface{}{
			"filename":    handler.Filename,
			"archive_size": handler.Size,
			"total_size":   0,
			"total_files":  0,
			"files":       []map[string]interface{}{},
		}
		fmt.Print("TESTING")

		zipReader, err := readZipFile(file)
		fmt.Print("TESTING")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, file := range zipReader.File {
			fileInfo := map[string]interface{}{
				"file_path": file.Name,
				"size":      file.FileInfo().Size(),
				"mimetype":  http.DetectContentType([]byte(file.FileInfo().Name())),
			}
			archiveInfo["total_size"] = archiveInfo["total_size"].(int64) + file.FileInfo().Size()
			archiveInfo["total_files"] = archiveInfo["total_files"].(int) + 1
			archiveInfo["files"] = append(archiveInfo["files"].([]map[string]interface{}), fileInfo)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%v", archiveInfo)
	})

	http.ListenAndServe(":8080", nil)
}

func readZipFile(file io.Reader) (*zip.Reader, error) {
	// Save the uploaded file temporarily
	tempFile, err := os.CreateTemp("", "uploaded-*.zip")
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return nil, err
	}

	

	fileInfo, err := tempFile.Stat()
	if err != nil {
		return nil, err
	}

	// Open the saved file as a ZIP archive
	tempFile.Seek(0, 0)
	zipReader, err := zip.NewReader(tempFile, fileInfo.Size())
	if err != nil {
		return nil, err
	}
	return zipReader, err
}

