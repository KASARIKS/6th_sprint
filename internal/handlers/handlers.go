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

	buf, err := getFileDataFromForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resFile, err := os.Create("res.txt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	io.WriteString(resFile, service.Convert(buf.String()))
	io.WriteString(w, service.Convert(buf.String()))
}

func getFileDataFromForm(r *http.Request) (bytes.Buffer, error) {
	file, _, err := r.FormFile("myFile")
	if err != nil {
		return *bytes.NewBuffer([]byte{}), err
	}
	defer file.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		return *bytes.NewBuffer([]byte{}), err
	}

	return buf, nil
}
