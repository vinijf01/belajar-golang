package belajargolangweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadFrom(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "upload_form.gohtml", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// r.ParseMultipartForm(32 << 20)
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	filesDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(filesDestination, file)
	if err != nil {
		panic(err)
	}

	//untuk mengambil value yang bukan file
	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload_success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadFrom)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/wallpaperflare.com_wallpaper.jpg
var uploadFileTest []byte

func TestUpload(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Vini Jumatul Fitri")
	file, _ := writer.CreateFormFile("file", "CONTOH.png")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	// request.Header.Set("Content-Type", "multipart/form-data")
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	response, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(response))

}
