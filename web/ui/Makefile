# Copyright 2021 Google LLC
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
ifeq ($(VERBOSE), 1)
	VERB =
endif

PORT = 5000

node-install:
	$(VERB) yarn install

node_modules: node-install

node-build: node_modules
	$(VERB) yarn run build

py-serve:
	$(VERB) python -m http.server --bind 127.0.0.1 --directory build $(PORT)