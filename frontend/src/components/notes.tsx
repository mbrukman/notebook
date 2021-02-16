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

class Notes extends Component {
  constructor() {
    super();
    this.showAddNote = this.showAddNote.bind(this);
    this.onChange = this.onChange.bind(this);
    this.saveNote = this.saveNote.bind(this);
    this.state = { adding: false, notes: [], newNote: { title: '', body: '' } };
  }

  showAddNote() {
    this.setState({ adding: true });
  }

  onChange(e) {
    this.setState((state) => {
      state.newNote[e.target.name] = e.target.value;
      return state;
    });
  }

  saveNote(e) {
    this.setState((state) => {
      let addNote = {
        title: state.newNote.title,
        body: state.newNote.body,
        id: Date.now(),
      };
      return {
        notes: [addNote].concat(state.notes),
        newNote: {
          title: '',
          body: '',
        },
        adding: false,
      };
    });
    e.preventDefault();
  }

  render(props, state) {
    return (
        <>
          <div>Notes</div>
          {
            !state.adding
            ? <button onclick={this.showAddNote}>Add note</button>
            : <form onsubmit={this.saveNote}>
                <b>Title</b><input type="text" name="title" onchange={this.onChange} /><br/>
                <b>Body</b><input type="textarea" name="body" onchange={this.onChange} />
                <button onclick={this.saveNote}>Add note</button>
              </form>
          }
          {
            state.notes.map((note) => (
                <>
                  <Note
                      title={note.title}
                      body={note.body}
                      key={note.id}>
                  </Note>
                  <hr/>
                </>
            ))
          }
        </>
    );
  }
}

export default Notes;
