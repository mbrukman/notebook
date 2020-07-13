// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/robfig/soy"
	"github.com/robfig/soy/data"
	"github.com/robfig/soy/soyhtml"
)

var (
	cwd, _  = os.Getwd()
	webRoot = flag.String("web-root", cwd, "Root of the web file tree.")
	host    = flag.String("host", "127.0.0.1", "By default, the server is only accessible via localhost. "+
		"Set to 0.0.0.0 or empty string to open to all.")
	port               = flag.String("port", getEnvWithDefault("PORT", "8080"), "Port to listen on; $PORT env var overrides default value.")
	tofu *soyhtml.Tofu = nil
)

func getEnvWithDefault(varName, defaultValue string) string {
	if value := os.Getenv(varName); value != "" {
		return value
	}
	return defaultValue
}

func dispatchHandler(rw http.ResponseWriter, req *http.Request) {
	if strings.HasSuffix(req.URL.Path, ".css") || strings.HasSuffix(req.URL.Path, ".js") {
		staticFileHandler(rw, req)
	} else if req.URL.Path == "/" {
		indexHandler(rw, req)
	} else {
		http.Error(rw, fmt.Sprintf("Not found: %s", req.URL.Path), 404)
	}
}

func indexHandler(rw http.ResponseWriter, req *http.Request) {
	log.Printf("Access [index]: %s %s", req.Method, req.URL.Path)
	var buf bytes.Buffer
	var m = make(data.Map)
	if err := tofu.Render(&buf, "home.index", m); err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}
	io.Copy(rw, &buf)
}

func staticFileHandler(rw http.ResponseWriter, req *http.Request) {
	log.Printf("Access [asset]: %s %s", req.Method, req.URL.Path)
	http.ServeFile(rw, req, path.Join(*webRoot, req.URL.Path))
}

func main() {
	flag.Parse()

	compiledTofu, err := soy.NewBundle().WatchFiles(true).AddTemplateDir(*webRoot).CompileToTofu()
	if err != nil {
		log.Fatal("Error compiling Soy templates: ", err.Error())
	}
	tofu = compiledTofu

	http.HandleFunc("/", dispatchHandler)

	hostPort := fmt.Sprintf("%s:%s", *host, *port)
	log.Printf("Listening on %s", hostPort)
	log.Fatal(http.ListenAndServe(hostPort, nil))
}
