package todolib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetVersion get app inform
func GetVersion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := DbQuery("version")
	b := DbQuery("name")
	c, _ := strconv.Atoi(DbQuery("updateTime"))

	fmt.Fprintln(w, JSONEnc(a, b, c))
}

// UploadFile ....
func UploadFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseMultipartForm(0 << 512)
	file, handler, err := r.FormFile("filename")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	//fmt.Printf("File Size: %+v\n", handler.Size)
	//fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := os.Create("./images/"+handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	MakeLog("UploadFile success")
}
