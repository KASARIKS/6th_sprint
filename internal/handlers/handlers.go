package handlers

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

type fileComponents struct {
	buffer     bytes.Buffer
	resultFile *os.File
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
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

	resFile, err := os.Create(time.Now().Format("02_01_2006_15_04_05") + ".txt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fileComps := fileComponents{
		buffer:     buf,
		resultFile: resFile,
	}

	err = writeData(fileComps, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

func writeData(fileComps fileComponents, w http.ResponseWriter) error {
	_, err := io.WriteString(fileComps.resultFile, service.Convert(fileComps.buffer.String()))
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, service.Convert(fileComps.buffer.String()))
	if err != nil {
		return err
	}

	return nil
}
