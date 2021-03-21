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

type InMemoryDatabase struct {
	storage []Note
}

func (inmem *InMemoryDatabase) ListNotes() []Note {
	return inmem.storage
}

func (inmem *InMemoryDatabase) CreateNote(note PartialNote) (Note, error) {
	newNote := Note{
		Title: note.Title,
		Body:  note.Body,
		Id:    fmt.Sprintf("%d", time.Now().UnixNano()/1000),
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

func (inmem *InMemoryDatabase) DeleteNote(id string) error {
	for idx, note := range inmem.storage {
		if note.Id == id {
			inmem.storage = append(inmem.storage[:idx], inmem.storage[idx+1:]...)
			return nil
		}
	}
	return notFoundError{message: fmt.Sprintf("Note id not found: %s", id)}
}

func CreateInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{}
}
