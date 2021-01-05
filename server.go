package main

import (
	"1_todo/todolib"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	todolib.MakeLog("Server started...")
	defer todolib.FpLog.Close() //logfile close

	router := httprouter.New()

	router.GET("/", todolib.Index)
	router.GET("/hello/:name", todolib.Hello)
	router.GET("/get/db/version", todolib.GetVersion)

	router.POST("/post/db/version", todolib.InsertVersion)
	router.POST("/signup", todolib.SignUp)
	router.POST("/signin", todolib.SignIn)

	router.PUT("/upload", todolib.UploadFile)
	router.DELETE("/delete/db/version", todolib.DeleteVersion)

	router.GET("/files", todolib.DownloadFile)

	err := http.ListenAndServe(":19124", router)
	if err != nil {
		todolib.MakeLog("ListenAndServe fail")
		panic(err)
	}
}
