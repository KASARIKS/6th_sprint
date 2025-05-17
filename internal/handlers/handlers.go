package handlers

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := readPageFile()
	if err != nil {
		http.Error(w, "Internal problems! Try later or contact support.", http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func readPageFile() ([]byte, error) {
	page, err := os.ReadFile("index.html")
	return page, err
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong method!", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	var buf bytes.Buffer
	io.Copy(&buf, file)

	io.WriteString(w, service.Convert(buf.String()))
}
