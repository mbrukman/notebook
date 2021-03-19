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

package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Id    string `json:"id"`
}

type ListNotesResponse struct {
	Notes []Note `json:"notes"`
}

type CreateNoteRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type CreateNoteResponse struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Id    string `json:"id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func listNotes(rw http.ResponseWriter, req *http.Request) {
	resp := &ListNotesResponse{
		Notes: make([]Note, 0),
	}
	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error serializing ListNotesResponse to JSON: %v\n", err)
	}
	fmt.Fprintf(rw, "%s", data)
}

func createNote(rw http.ResponseWriter, req *http.Request) {
	body, postErr := ioutil.ReadAll(req.Body)
	if postErr != nil {
		log.Printf("Error reading POST body for Create: %v\n", postErr)
		return
	}
	log.Printf("CreateNoteRequest (JSON): %s\n", body)
	createReq := &CreateNoteRequest{}
	jsonErr := json.Unmarshal(body, &createReq)
	if jsonErr != nil {
		log.Printf("Error parsing JSON for CreateNoteRequest: %v\n", jsonErr)
	}
	createResp := &CreateNoteResponse{
		Title: createReq.Title,
		Body:  createReq.Body,
		Id:    fmt.Sprintf("%d", time.Now().UnixNano()/1000),
	}
	data, err := json.Marshal(createResp)
	if err != nil {
		log.Printf("Error serializing CreateNoteResponse to JSON: %v\n", err)
	}
	log.Printf("CreateNoteResponse (JSON): %s\n", data)
	fmt.Fprintf(rw, "%s", data)
}

func handleError(rw http.ResponseWriter, req *http.Request) {
	resp := &ErrorResponse{
		Error: fmt.Sprintf("Unrecognized API path: %s", req.URL.Path),
	}
	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error serializing ErrorResponse to JSON: %v\n", err)
	}
	fmt.Fprintf(rw, "%s", data)
}

func HandleRequest(rw http.ResponseWriter, req *http.Request) {
	log.Printf("Request [API]: %s %s", req.Method, req.URL.Path)
	if req.URL.Path == "/api/v1/notes" {
		listNotes(rw, req)
	} else if req.URL.Path == "/api/v1/notes/create" {
		createNote(rw, req)
	} else {
		handleError(rw, req)
	}
}
