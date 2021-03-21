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

package db

import (
	"fmt"
	"time"
)

// InMemoryDatabase is a simgple in-memory storage system for local
// prototyping, development, and testing. As it is not durable, it's not
// intended to be used in production.
type InMemoryDatabase struct {
	storage []Note
}

// ListNotes returns all notes stored in the database.
func (inmem *InMemoryDatabase) ListNotes() []Note {
	return inmem.storage
}

// CreateNote creates a note in the in-memory database.
func (inmem *InMemoryDatabase) CreateNote(note PartialNote) (Note, error) {
	newNote := Note{
		Title: note.Title,
		Body:  note.Body,
		ID:    fmt.Sprintf("%d", time.Now().UnixNano()/1000),
	}
	inmem.storage = append(inmem.storage, newNote)
	return newNote, nil
}

type notFoundError struct {
	message string
}

func (notFound notFoundError) Error() string {
	return notFound.message
}

// DeleteNote deletes the note with the given id from the in-memory database.
// If the note with such an ID does not exist, or if the deletion request
// failed for any other reason, an error is returned.
func (inmem *InMemoryDatabase) DeleteNote(id string) error {
	for idx, note := range inmem.storage {
		if note.ID == id {
			inmem.storage = append(inmem.storage[:idx], inmem.storage[idx+1:]...)
			return nil
		}
	}
	return notFoundError{message: fmt.Sprintf("Note id not found: %s", id)}
}

// CreateInMemoryDatabase returns a new instance of the in-memory database.
func CreateInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{}
}
