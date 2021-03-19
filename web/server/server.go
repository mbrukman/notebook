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
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/mbrukman/notebook/web/server/api"
)

var (
	cwd, _  = os.Getwd()
	webRoot = flag.String("web-root", cwd, "Root of the web file tree.")
	host    = flag.String("host", "127.0.0.1", "By default, the server is only accessible via localhost. "+
		"Set to 0.0.0.0 or empty string to open to all.")
	port = flag.String("port", getEnvWithDefault("PORT", "8080"), "Port to listen on; $PORT env var overrides default value.")
)

func getEnvWithDefault(varName, defaultValue string) string {
	if value := os.Getenv(varName); value != "" {
		return value
	}
	return defaultValue
}

func handleStaticFile(rw http.ResponseWriter, req *http.Request) {
	log.Printf("Request [static]: %s %s", req.Method, req.URL.Path)
	http.ServeFile(rw, req, path.Join(*webRoot, req.URL.Path))
}

func dispatchHandler(rw http.ResponseWriter, req *http.Request) {
	if strings.HasPrefix(req.URL.Path, "/api/") {
		api.HandleRequest(rw, req)
	} else {
		handleStaticFile(rw, req)
	}
}

func main() {
	flag.Parse()

	http.HandleFunc("/", dispatchHandler)

	hostPort := fmt.Sprintf("%s:%s", *host, *port)
	log.Printf("Listening on %s", hostPort)
	log.Fatal(http.ListenAndServe(hostPort, nil))
}
