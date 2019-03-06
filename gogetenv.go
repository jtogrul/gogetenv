package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type envInfo struct {
	Path   string              `json:"url_path"`
	Params map[string][]string `json:"url_params"`
	Env    []string            `json:"environment"`
}

func serveEnvInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	envInfo := envInfo{
		Path:   r.URL.Path,
		Params: r.Form,
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
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
}
