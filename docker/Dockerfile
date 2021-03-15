# Copyright 2021 Google Inc.
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

FROM node:15.11.0-alpine3.13

# To be able to run inside the Docker container as a non-root user and still
# have read/write access to the file system, we assign the smae UID to the
# non-root user inside the container as we are using outside of the container.
# For this reason, the build step passes in the UID of the external user, and
# this is set as the UID of the user inside the container.
ARG UID

# Install the `usermod` tool we will need to use below.
RUN apk --no-cache add shadow

# Change the UID of the pre-existing `node` user to match that of the user
# building this container, such that:
#
# (a) the user inside the container can access files owned by the user outside
#     of the container, and
# (b) files created by the user inside the container are owned by the user
#     outside of the container.
#
# This allows seamless usage of the container, without using root permissions.
RUN usermod -u $UID node
USER node

VOLUME /opt/notebook
WORKDIR /opt/notebook

# `preact watch` uses port 8080 by default.
EXPOSE 8080/tcp
