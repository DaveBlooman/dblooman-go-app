package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultProdStack...)
	router, err := rest.MakeRouter(
		rest.Get("/status", GetStatus),
		rest.Get("/", GetRoot),
		rest.Get("/meta", GetMeta),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func GetStatus(w rest.ResponseWriter, r *rest.Request) {
	response := map[string]string{"status": "all good mate"}
	w.WriteJson(response)
}

func GetMeta(w rest.ResponseWriter, r *rest.Request) {
	response, _ := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	output := map[string]string{"status": string(contents)}
	w.WriteJson(output)
}

func GetRoot(w rest.ResponseWriter, r *rest.Request) {
	output := map[string]string{"status": "hey hey hey"}
	w.WriteJson(output)
}
