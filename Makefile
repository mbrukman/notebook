# Copyright 2017 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

VERB = @
ifeq ($(VERBOSE),1)
	VERB =
endif

gofmt:
	$(VERB) find . -name '*.go' | xargs -I {} gofmt -s -w {}

gofmt_test:
	$(VERB) echo "Running 'go fmt' test ..."
	$(VERB) ./gofmt_test.sh

go_mod_tidy_test:
	$(VERB) echo "Running 'go mod tidy' test ..."
	$(VERB) ./go_mod_tidy_test.sh

go-test:
	$(VERB) go test ./...

govet:
	$(VERB) echo "Running 'go vet' ..."
	$(VERB) ./go_vet_test.sh

go-update-workspace:
	$(VERB) bazel run //:gazelle -- update-repos -from_file=go.mod

go-update-build:
	$(VERB) bazel run //:gazelle -- -build_file_name BUILD.bazel

test: gofmt_test go_mod_tidy_test go-test govet
