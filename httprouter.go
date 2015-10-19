package main

import (
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type jsonReq struct {
	Name string
}

type jsonRes struct {
	Greeting string
}


func helloGet(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("Abdul"))
}

func helloPost(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(req.Body)
	var reqJson jsonReq
	err := decoder.Decode(&reqJson)
	if err != nil {
		panic("Error!! - JSON Decoding")
	}

	Res := jsonRes{Greeting: "Hello, " + reqJson.Name+"!"}
	json.NewEncoder(rw).Encode(Res)

}
func main() {
	mux := httprouter.New()
	mux.GET("/hello/:Abdul", helloGet)
	mux.POST("/hello/", helloPost)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
