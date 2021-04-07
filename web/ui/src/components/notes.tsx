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

import { Component } from 'preact';
import Note from './note';

interface NotesState {
  adding: boolean;
  notes: Array<{
    title: string;
    body: string;
    id: string;
  }>;
  newNote: {
    title: string;
    body: string;
  }
}

interface CreateNoteResponse {
  title: string;
  body: string;
  id: string;
}

interface ListNotesResponse {
  notes: Array<{
    title: string;
    body: string;
    id: string;
  }>;
}

class Notes extends Component {
  constructor() {
    super();
    this.saveNoteClientSide = this.saveNoteClientSide.bind(this);
    this.saveNoteServerSide = this.saveNoteServerSide.bind(this);
    this.state = {
      adding: false,
      notes: [],
      newNote: {
        title: '',
        body: '',
      },
    } as NotesState;
  }

  componentDidMount() {
    fetch('/api/v1/notes')
    .then((response: Response) => {
      if (response.status !== 200) {
        response.text().then((errText: string) => {
          // TODO(mbrukman): display error in the UI as well as the console.
          console.log('Error (' + response.status + '): ' + errText);
        });
        return;
      }
      response.json().then((listNotes: ListNotesResponse) => {
        this.setState((s: NotesState) => ({
          adding: false,
          notes: listNotes.notes,
          newNote: {
            title: '',
            body: '',
          }
        }) as NotesState);
      });
    })
    .catch(console.log);
  }

  saveNoteServerSide(e: Event) {
    const toAddNote = {
      title: this.state.newNote.title,
      body: this.state.newNote.body,
    };
    fetch('/api/v1/notes/create', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(toAddNote),
    })
    .then((response: Response) => {
      if (response.status !== 200) {
        response.text().then((errText: string) => {
          // TODO(mbrukman): display error in the UI as well as the console.
          console.log('Error (' + response.status + '): ' + errText);
        });
        return;
      }
      response.json().then((addedNote: CreateNoteResponse) => {
        this.setState({
          notes: [addedNote].concat(this.state.notes),
          newNote: {
            title: '',
            body: '',
          },
          adding: false,
        });
      })
    })
    .catch(console.log);

    e.preventDefault();
    return false;
  }

  saveNoteClientSide(e: Event) {
    const addNote = {
      title: this.state.newNote.title,
      body: this.state.newNote.body,
      id: Date.now().toString(),
    };
    this.setState({
      notes: [addNote].concat(this.state.notes),
      newNote: {
        title: '',
        body: '',
      },
      adding: false,
    });
    e.preventDefault();
    return false;
  }

  render(props, state: NotesState) {
    const onChange = (e: Event) => {
      this.setState((state: NotesState) => {
        state.newNote[e.target.name] = e.target.value;
        return state;
      });
    };

    const showAddNote = (e: Event) => {
      this.setState({
        adding: true,
        newNote: {
          title: '',
          body: '',
        },
      });
    };

    const cancelSaveNote = (e: Event) => {
      this.setState({ adding: false });
      return false;
    };

    return (
        <>
          <div>Notes</div>
          {
            !state.adding
            ? <button onclick={showAddNote}>Add note</button>
            : <div>
                <b>Title</b><input type="text" name="title" onchange={onChange} /><br/>
                <b>Body</b><textarea name="body" onchange={onChange} />
                <button onclick={this.saveNoteServerSide}>Save note</button>
                <button onclick={cancelSaveNote}>Cancel</button>
              </div>
          }
          {
            state.notes.map((note) => (
                <Note
                    title={note.title}
                    body={note.body}
                    key={note.id}>
                </Note>
            ))
          }
        </>
    );
  }
}

export default Notes;
