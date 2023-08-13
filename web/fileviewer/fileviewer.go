package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gomarkdown/markdown"
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

func stringHasOneOfSuffixes(str string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(str, suffix) {
			return true
		}
	}
	return false
}

type DocHandler struct {
	webRoot string
}

func serveFile(rw http.ResponseWriter, mimeType string, fileContents []byte) {
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", fmt.Sprintf("%s; charset=UTF-8", mimeType))
	rw.Write([]byte(fileContents))
}

func (handler *DocHandler) DispatchHandler(rw http.ResponseWriter, req *http.Request) {
	log.Printf("Request: %s", req.URL.Path)
	var urlPath string = req.URL.Path
	// Remove any occurrences of `..` in the path to avoid escaping the web root.
	urlPath = strings.ReplaceAll(urlPath, "..", "")
	// Replace all multi-sequences of `/` with a single slash.
	urlPath = strings.ReplaceAll(urlPath, "//", "/")
	var fsPath string = path.Join(handler.webRoot, "/./", urlPath)

	log.Printf("Local path: %s", fsPath)

	fileInfo, err := os.Stat(fsPath)
	if err != nil && os.IsNotExist(err) {
		// File does not exist.
		rw.WriteHeader(http.StatusNotFound)
		rw.Header().Set("Content-Type", "text/html; charset=UTF-8")
		rw.Write([]byte("<!DOCTYPE html>\n"))
		rw.Write([]byte("<html>"))
		rw.Write([]byte("<body>"))
		rw.Write([]byte(fmt.Sprintf("<h1>Error 404: <code>%s</code> not found", urlPath)))
		rw.Write([]byte("</body>"))
		rw.Write([]byte("</html>"))

		log.Printf("Path not found: %s", fsPath)
		return
	}

	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(fsPath)
		if err != nil {
			log.Printf("Error listing directory: %s", fsPath)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "text/html; charset=UTF-8")
		rw.Write([]byte("<!DOCTYPE html>\n"))
		rw.Write([]byte("<html>"))
		rw.Write([]byte("<body>"))
		rw.Write([]byte(fmt.Sprintf("<h1>Directory listing: %s</h1>", urlPath)))
		rw.Write([]byte("<ul>"))
		for _, file := range files {
			// Skip internal files, e.g., `.git` directory, `.gitignore`, other dotfiles, etc.
			if strings.HasPrefix(file.Name(), ".") {
				continue
			}
			var listItem string
			if urlPath == "/" {
				listItem = fmt.Sprintf("<li><a href='%s'>%s</li>\n", file.Name(), file.Name())
			} else {
				listItem = fmt.Sprintf("<li><a href='%s/%s'>%s</li>\n", urlPath, file.Name(), file.Name())
			}
			rw.Write([]byte(listItem))
		}
		rw.Write([]byte("</ul>"))
		rw.Write([]byte("</body>"))
		rw.Write([]byte("</html>"))
		return
	}

	fileContents, err := ioutil.ReadFile(fsPath)
	if err != nil {
		log.Printf("Error reading file (%s): %s", fsPath, err)
		return
	}

	if strings.HasSuffix(fsPath, ".html") {
		serveFile(rw, "text/html", fileContents)
	} else if strings.HasSuffix(fsPath, ".js") {
		serveFile(rw, "text/javascript", fileContents)
	} else if strings.HasSuffix(fsPath, ".ts") {
		serveFile(rw, "text/typescript", fileContents)
	} else if strings.HasSuffix(fsPath, ".css") {
		serveFile(rw, "text/css", fileContents)
	} else if strings.HasSuffix(fsPath, ".md") {
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "text/html; charset=UTF-8")
		rw.Write([]byte(`<!doctype html>
<html>
<head>
  <style>
    code {
      background-color: #efefef;
      margin: 3px;
      padding: 3px;
    }
  </style>
</head>
<body>`))
		rw.Write(markdown.ToHTML(fileContents, nil, nil))
		rw.Write([]byte(`</body>
</html>`))
	} else if stringHasOneOfSuffixes(fsPath, []string{".txt", ".text", ".json", ".sh"}) {
		rw.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		rw.Write(fileContents)
	} else {
		rw.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		rw.Write([]byte("Unrecognized file content type or suffix."))
		log.Printf("Unrecognized file content type or suffix: %s", fsPath)
	}
}

func NewDocHandler(webRoot string) *DocHandler {
	return &DocHandler{webRoot: webRoot}
}

func main() {
	flag.Parse()

	handler := NewDocHandler(*webRoot)
	http.HandleFunc("/", handler.DispatchHandler)

	hostPort := fmt.Sprintf("%s:%s", *host, *port)
	log.Printf("Listening on http://%s", hostPort)
	log.Printf("Serving from %s", *webRoot)
	log.Fatal(http.ListenAndServe(hostPort, nil))
}
