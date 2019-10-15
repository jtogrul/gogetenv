package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type envInfo struct {
	Host   string              `json:"host"`
	Path   string              `json:"path"`
	Params map[string][]string `json:"url_params"`
	Header map[string][]string `json:"header"`
	Env    []string            `json:"environment"`
}

func serveEnvInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	envInfo := envInfo{
		Host:   r.Host,
		Path:   r.URL.Path,
		Params: r.Form,
		Header: r.Header,
		Env:    os.Environ(),
	}

	response, err := json.MarshalIndent(envInfo, "", "\t")

	if err != nil {
		fmt.Fprintln(w, "Error generating JSON response")
		log.Println("Error generating JSON response: ", err)
	} else {
		fmt.Fprintln(w, string(response))
	}
}

func main() {
	http.HandleFunc("/", serveEnvInfo)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
}
