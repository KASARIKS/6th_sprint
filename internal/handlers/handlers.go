package handlers

import (
	"io"
	"net/http"
	"os"
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

	_, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	io.WriteString(w, header.Filename)
}
