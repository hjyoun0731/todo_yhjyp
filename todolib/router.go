package todolib

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// Index main page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// Hello hello page
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

// GetVersion get app inform
func GetVersion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := DbQuery("version", "version")
	b := DbQuery("name", "version")
	c, _ := strconv.Atoi(DbQuery("updateTime", "version"))

	mem := version{a, b, c}

	fmt.Fprintln(w, JSONEnc(mem))
	MakeLog("GetVersion success")
}

func InsertVersion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024))
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
	newVersion := JSONDecVer(body)
	DbVerInsert(newVersion)

	MakeLog("InsertVersion success")
}

// UploadFile ....
func UploadFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseMultipartForm(0 << 100)
	file, handler, err := r.FormFile("filename")
	if err != nil {
		fmt.Println(err)

		fmt.Fprint(w,r.Header)
		fmt.Fprint(w,r.Body)

		w.Write([]byte("\nFail."))
		return
	}
	defer file.Close()

	upFile, err := os.Create("./files/" + handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer upFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	upFile.Write(fileBytes)
	MakeLog("UploadFile success")
}

func DownloadFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := "./files/가르쳐줘 코딩 소녀.apk"

	content, err := ioutil.ReadFile(path)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		panic(err)
	}

	w.Header().Add("Content-Type", "multipart/form-data")
	w.Write(content)
}

func DeleteVersion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := DbQuery("version", "version")
	DbDelete(a)
	MakeLog("DeleteVersion success")
}

func SignUp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	pw := ps.ByName("password")
	check := ps.ByName("check")

	dup := DbQuery("userid", "account")
	if dup == id {
		MakeLog("id duplicated")
		return
	}

	if pw == check {
		mem := account{id, pw}
		DbAcctInsert(mem)
		MakeLog("sign up success")
	}
}

func SignIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if DbQuery("userid", "account") == ps.ByName("id") && DbQuery("password", "account") == ps.ByName("password") {
		fmt.Fprint(w, "SignIn OK")
		MakeLog("sign in success")
	} else {
		fmt.Fprint(w, "SignIn fail")
		MakeLog("sign in fail(userid or password is wrong)")
	}
}
