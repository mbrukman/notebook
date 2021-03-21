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

// PartialNote represents a note as may be sent by a user when creating a note,
// e.g., missing an Id field.
type PartialNote struct {
	Title string
	Body  string
}

// Note represents a complete note, e.g., as stored in the database.
type Note struct {
	Title string
	Body  string
	ID    string
}

// Database represents a generic abstract storage system, with several
// potential implementations.
type Database interface {
	ListNotes() []Note
	CreateNote(note PartialNote) (Note, error)
	DeleteNote(id string) error
}
