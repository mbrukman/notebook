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

ifeq ($(VERBOSE),1)
	VERB =
endif

GO_SRCS := $(shell find . -name '*.go')
PORT = 7000

proxy: $(GO_SRCS)
	$(VERB) go build .

run: proxy
	$(VERB) ./proxy -config config.yaml -port $(PORT)
