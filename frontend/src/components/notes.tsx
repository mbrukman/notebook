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
    this.state = { notes: [{ title: 'Note 1', body: 'Contents', id: '12345' }] };
  }

  render(props, state) {
    return (
        <>
          <div>Notes</div>
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
