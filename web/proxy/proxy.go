// Copyright 2021 Google LLC
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
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/ghodss/yaml"
)

var (
	host = flag.String("host", "127.0.0.1", "Host to listen on; 127.0.0.1 only allows connections "+
		"from localhost; set to 0.0.0.0 or empty to allow external connections.")
	port       = flag.Int("port", 9000, "Port to listen on")
	configFile = flag.String("config", "", "Config file, either YAML (*.yaml, *.yml) or JSON (*.json) format.")
)

// Route represents a single pair (URL path prefix, target).
type Route struct {
	Path   string `yaml:"path" json:"path"`
	Target string `yaml:"target" json:"target"`
}

// ProxyConfig represents the set of routes for the reverse proxy.
type ProxyConfig struct {
	Routes []Route `yaml:"routes" json:"routes"`
}

// ConfigError is returned while parsing or loading an invalid config.
type ConfigError struct {
	Message string
}

func (ce *ConfigError) Error() string {
	return ce.Message
}

// ParseConfig loads the config stored in the provided file and returns a
// config object.
func ParseConfig(filename string) (*ProxyConfig, error) {
	if filename == "" {
		return nil, &ConfigError{
			Message: "Config filename must be specified (provided empty string)",
		}
	}

	data, readErr := ioutil.ReadFile(filename)
	if readErr != nil {
		return nil, &ConfigError{
			Message: fmt.Sprintf("Error reading file %s: %v", filename, readErr),
		}
	}

	config := &ProxyConfig{}
	if strings.HasSuffix(*configFile, "yaml") || strings.HasSuffix(*configFile, ".yml") {
		yamlErr := yaml.Unmarshal(data, config)
		if yamlErr != nil {
			return nil, &ConfigError{
				Message: fmt.Sprintf("Error parsing YAML: %v", yamlErr),
			}
		}
	} else if strings.HasSuffix(*configFile, ".json") {
		jsonErr := json.Unmarshal(data, config)
		if jsonErr != nil {
			return nil, &ConfigError{
				Message: fmt.Sprintf("Error parsing JSON: %v", jsonErr),
			}
		}
	} else {
		return nil, &ConfigError{
			Message: fmt.Sprintf("Unrecognized file suffix (only YAML and JSON are supported): %s\n", *config),
		}
	}
	return config, nil
}

// Proxy represents the runtime state of the reverse proxy.
type Proxy struct {
	Backends map[string]*httputil.ReverseProxy
}

func (proxy *Proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	for path, backend := range proxy.Backends {
		if strings.HasPrefix(req.URL.Path, path) {
			log.Printf("%s %s", req.Method, req.URL.Path)
			backend.ServeHTTP(rw, req)
			return
		}
	}
	log.Printf("Error: unhandled request (no backend mapping): %s %s\n", req.Method, req.URL.Path)
}

func main() {
	flag.Parse()

	config, configErr := ParseConfig(*configFile)
	if configErr != nil {
		log.Printf("%v\n", configErr)
		os.Exit(1)
	}

	backends := make(map[string]*httputil.ReverseProxy)
	for _, route := range config.Routes {
		target, err := url.Parse(route.Target)
		if err != nil {
			log.Fatalf("Error parsing target URL %s: %v\n", route.Target, err)
		}
		log.Printf("Redirecting %s -> %s\n", route.Path, route.Target)
		backends[route.Path] = httputil.NewSingleHostReverseProxy(target)
	}

	hostPort := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Listening on http://%s", hostPort)

	proxy := &Proxy{
		Backends: backends,
	}

	http.Handle("/", proxy)
	log.Fatal(http.ListenAndServe(hostPort, nil))
}
