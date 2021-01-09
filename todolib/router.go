package todolib

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Index main page
func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_,err := fmt.Fprint(w, "Welcome!\n")
	if err != nil {
		MakeLog("router.Index error")
	}
}

// Hello hello page
func Hello(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	_, err := fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
	if err != nil {
		MakeLog("router.hello error")
	}
}

// GetVersion get app newest version information
func GetVersion(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	a := DbQuery("version", "version")
	b := DbQuery("name", "version")
	c := DbQuery("updateTime", "version")

	mem := version{a, b, c}

	_,err := fmt.Fprintln(w, JSONEnc(mem))
	if err != nil {
		MakeLog("GetVersion write error")
	}
	MakeLog("GetVersion success")
}

func InsertVersion(_ http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
func UploadFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseMultipartForm(0 << 100)
	if err != nil {
		MakeLog("UploadFile ParseMultipartForm fail")
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(404)
		//w.Write([]byte(http.StatusText(404)))

		MakeLog("UploadFile Fail.")
		return
	}
	defer file.Close()

	latestVer := version{r.FormValue("version"), r.FormValue("name"), r.FormValue("updateTime") }
	DbVerInsert(latestVer)

	upFile, err := os.Create("./files/" + handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer upFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	_,err = upFile.Write(fileBytes)
	if err != nil {
		MakeLog("UploadFile fileBytes write fail")
	}
	MakeLog("UploadFile success")
}

func DeleteVersion(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	a := DbQuery("version", "version")
	DbDelete(a)
	MakeLog("DeleteVersion success")
}

func SignUp(_ http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id,pw,check := r.FormValue("id"),r.FormValue("password"),r.FormValue("check")

	dup := DbQuery("userid", "account")
	if dup == id {
		MakeLog("id duplicated")
		return
	}

	if pw == check {
		mem := account{id, pw}
		DbAcctInsert(mem)
		MakeLog("sign up success")
	} else {
		MakeLog("Sign up fail. pw1, pw2 is different")
	}
}

func SignIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if DbQuery("userid", "account") == r.FormValue("id") && DbQuery("password", "account") == r.FormValue("password") {
		_,err := fmt.Fprint(w, "SignIn OK")
		if err != nil {
			panic(err)
		}
		MakeLog("sign in success")
	} else {
		_,err := fmt.Fprint(w, "SignIn fail")
		if err != nil {
			panic(err)
		}
		MakeLog("sign in fail(userid or password is wrong)")
	}
}

func GetPath(w http.ResponseWriter, _ *http.Request, ps httprouter.Params){
	var list []string
	var dVer string

	ver := ps.ByName("version")
	if ver == "LTS" {
		dVer = "LTS.apk"
	} else {
		dVer = DbQuery("version","version")+".apk"
	}

	files, err := ioutil.ReadDir("./files")
	if err != nil {
		MakeLog("read dir fail")
		return
	}
	for _, filename := range files {
		list = append(list, filename.Name())
	}

	for _, f := range list {
		fVer := strings.Split(f,"_v")

		if len(fVer) < 2{
			continue
		}

		if fVer[1] == dVer {
			_,err := fmt.Fprint(w, "http://hjyoun.ddns.net:19124/path/"+f)
			if err != nil {
				MakeLog("GetPath write path fail")
			}
		}
	}

	MakeLog("download server")
}
