package main

import (
	"1_todo/todolib"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Index ....
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// Hello .....
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	todolib.MakeLog("Server started...")
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	router.GET("/db/version", todolib.GetVersion)
	router.PUT("/upload/:id/:name", todolib.UploadFile)

	err := http.ListenAndServe(":19124", router)
	if err != nil {
		todolib.MakeLog("ListenAndServe fail")
		panic(err)
	}
	defer todolib.FpLog.Close()
}
