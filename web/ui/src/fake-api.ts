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

import { createServer, Model } from "miragejs";
import { CreateNoteRequest, CreateNoteResponse } from "../components/notes";

if (JSON.parse(process.env.FAKE_API)) {
  createServer({
    models: {
      notes: Model,
    },

    routes() {
      this.namespace = "api";

      let noteId = 1;
      this.post("/v1/notes/create", (schema, req: CreateNoteReqest) => {
        let note: CreateNoteResponse = {
          ...JSON.parse(req.requestBody),
          id: noteId++,
        };

        schema.notes.create(note);
        return note;
      });

      this.get("/v1/notes", (schema) => schema.notes.all());
    },
  });
}
