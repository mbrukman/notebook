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
	"path"
	"strings"

	"github.com/mbrukman/notebook/web/server/db"
)

// Note represents a singular note.
type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	ID    string `json:"id"`
}

// ListNotesResponse is the response type of the ListNotes RPC.
//
// There is no parallel ListNotesRequest at this time as the user sends no
// payload for this RPC.
type ListNotesResponse struct {
	Notes []Note `json:"notes"`
}

// CreateNoteRequest is the request type of the CreateNote RPC.
type CreateNoteRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// CreateNoteResponse is the response type of the CreateNote RPC.
type CreateNoteResponse struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	ID    string `json:"id"`
}

// ErrorResponse is a generic error response which may be used in several RPCs.
type ErrorResponse struct {
	Error string `json:"error"`
}

func convertDatabaseNotesToApiNotes(notes []db.Note) []Note {
	result := make([]Note, 0)
	for _, note := range notes {
		result = append(result, Note{
			Title: note.Title,
			Body:  note.Body,
			ID:    note.ID,
		})
	}
	return result
}

func (apiHandler *ApiHandler) listNotes(rw http.ResponseWriter, req *http.Request) {
	notes := (*apiHandler.storage).ListNotes()
	resp := &ListNotesResponse{
		Notes: convertDatabaseNotesToApiNotes(notes),
	}
	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error serializing ListNotesResponse to JSON: %v\n", err)
	}
	fmt.Fprintf(rw, "%s", data)
}

func (apiHandler *ApiHandler) createNote(rw http.ResponseWriter, req *http.Request) {
	body, postErr := ioutil.ReadAll(req.Body)
	if postErr != nil {
		responseText := fmt.Sprintf("Error reading POST body for CreateNoteRequest: %v", postErr)
		log.Printf("%s\n", responseText)
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, responseText)
		return
	}
	log.Printf("CreateNoteRequest (JSON): %s\n", body)
	createReq := &CreateNoteRequest{}
	jsonErr := json.Unmarshal(body, &createReq)
	if jsonErr != nil {
		responseText := fmt.Sprintf("Error parsing JSON for CreateNoteRequest: %v", jsonErr)
		log.Printf("%s\n", responseText)
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, responseText)
		return
	}
	if createReq.Title == "" && createReq.Body == "" {
		responseText := fmt.Sprintf("Invalid CreateNoteRequest: `title` or `body` must be specified")
		log.Printf("%s\n", responseText)
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, responseText)
		return
	}
	storedNote, storeErr := (*apiHandler.storage).CreateNote(db.PartialNote{
		Title: createReq.Title,
		Body:  createReq.Body,
	})
	if storeErr != nil {
		responseText := fmt.Sprintf("Error calling CreateNoteRequest: %v", storeErr)
		log.Printf("%s\n", responseText)
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, responseText)
		return
	}
	createResp := &CreateNoteResponse{
		Title: storedNote.Title,
		Body:  storedNote.Body,
		ID:    storedNote.ID,
	}
	data, err := json.Marshal(createResp)
	if err != nil {
		responseText := fmt.Sprintf("Error serializing CreateNoteResponse to JSON: %v", err)
		log.Printf("%s\n", responseText)
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, responseText)
		return
	}
	log.Printf("CreateNoteResponse (JSON): %s\n", data)
	rw.WriteHeader(http.StatusCreated)
	fmt.Fprintf(rw, "%s", data)
}

func (apiHandler *ApiHandler) handleError(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(rw, "Unrecognized API path: %s", req.URL.Path)
}

// HandleRequest dispatches the incoming request based on URL path.
func (apiHandler *ApiHandler) handleRequest(rw http.ResponseWriter, req *http.Request) {
	log.Printf("Request [API]: %s %s", req.Method, req.URL.Path)
	if req.URL.Path == "/api/v1/notes" {
		apiHandler.listNotes(rw, req)
	} else if req.URL.Path == "/api/v1/notes/create" {
		apiHandler.createNote(rw, req)
	} else {
		apiHandler.handleError(rw, req)
	}
}

func (apiHandler *ApiHandler) handleStaticFile(rw http.ResponseWriter, req *http.Request) {
	log.Printf("Request [static]: %s %s", req.Method, req.URL.Path)
	http.ServeFile(rw, req, path.Join(apiHandler.webRoot, req.URL.Path))
}

func (apiHandler *ApiHandler) DispatchHandler(rw http.ResponseWriter, req *http.Request) {
	if strings.HasPrefix(req.URL.Path, "/api/") {
		apiHandler.handleRequest(rw, req)
	} else {
		apiHandler.handleStaticFile(rw, req)
	}
}

type ApiHandler struct {
	webRoot string
	storage *db.Database
}

func NewApiHandler(webRoot string, storage db.Database) *ApiHandler {
	return &ApiHandler{webRoot: webRoot, storage: &storage}
}
