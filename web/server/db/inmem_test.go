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
	"strings"
	"testing"
)

var (
	partialNote = PartialNote{
		Title: "title",
		Body:  "body",
	}
)

func matchField(t *testing.T, field string, expected string, actual string) {
	if expected != actual {
		t.Errorf("%s doesn't match: %s (expected) vs. %s (actual)", field, expected, actual)
	}
}

func matchPartialNoteAndNote(t *testing.T, expected PartialNote, actual Note) {
	matchField(t, "Title", expected.Title, actual.Title)
	matchField(t, "Body", expected.Body, actual.Body)
}

func matchNotes(t *testing.T, expected Note, actual Note) {
	matchField(t, "Title", expected.Title, actual.Title)
	matchField(t, "Body", expected.Body, actual.Body)
	matchField(t, "ID", expected.ID, actual.ID)
}

func expectEmpty(t *testing.T, database *InMemoryDatabase) {
	allNotes := database.ListNotes()
	if len(allNotes) != 0 {
		t.Errorf("Expected to have an empty set of notes; got %d", len(allNotes))
		return
	}
}

func TestSequence_Add_List_Delete_List(t *testing.T) {
	database := NewInMemoryDatabase()
	expectEmpty(t, database)

	// Add a note to the database.
	note, createErr := database.CreateNote(partialNote)
	if createErr != nil {
		t.Errorf("Error adding note: %v", createErr)
		return
	}
	matchPartialNoteAndNote(t, partialNote, note)

	// Ensure we now have just the 1 note we added above.
	allNotes := database.ListNotes()
	if len(allNotes) != 1 {
		t.Errorf("Expected to have exactly one note; got %d", len(allNotes))
		return
	}
	firstNote := allNotes[0]
	matchNotes(t, note, firstNote)

	// Delete the not we added earlier.
	deleteErr := database.DeleteNote(note.ID)
	if deleteErr != nil {
		t.Errorf("Error deleting note: %v", deleteErr)
		return
	}

	// The database should now be empty.
	expectEmpty(t, database)
}

func TestSequence_Add_List_DeleteMissing_List(t *testing.T) {
	database := NewInMemoryDatabase()
	expectEmpty(t, database)

	// Add a note to the database.
	note, createErr := database.CreateNote(partialNote)
	if createErr != nil {
		t.Errorf("Error adding note: %v", createErr)
		return
	}
	matchPartialNoteAndNote(t, partialNote, note)

	// Should now have 1 note after the addition.
	allNotes := database.ListNotes()
	if len(allNotes) != 1 {
		t.Errorf("Expected to have exactly one note; got %d", len(allNotes))
		return
	}
	firstNote := allNotes[0]
	matchNotes(t, note, firstNote)

	// Attempt to delete a missing note ID; expect it to fail.
	deleteErr := database.DeleteNote("this-id-does-not-exist")
	if deleteErr == nil {
		t.Errorf("Expected to fail deletion, since this note ID doesn't exist.")
		return
	}

	// Should still have 1 note after failed deletion.
	allNotes = database.ListNotes()
	if len(allNotes) != 1 {
		t.Errorf("Expected to have exactly one note; got %d", len(allNotes))
		return
	}
}

func TestDeleteFromEmptyDatabase(t *testing.T) {
	database := NewInMemoryDatabase()
	expectEmpty(t, database)

	// Attempt to delete a missing note ID; expect it to fail.
	deleteErr := database.DeleteNote("some-note-id")
	if deleteErr == nil {
		t.Errorf("Should have returned error for a missing note")
		return
	}
	if !strings.HasPrefix(deleteErr.Error(), "Note id not found: ") {
		t.Errorf("Expected error to have prefix `Note id not found:`, got: %s", deleteErr.Error())
		return
	}

	// The database should still be empty.
	expectEmpty(t, database)
}
