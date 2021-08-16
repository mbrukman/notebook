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

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "proxy_lib",
    srcs = ["proxy.go"],
    importpath = "github.com/mbrukman/notebook/web/proxy",
    visibility = ["//visibility:private"],
    deps = ["@com_github_ghodss_yaml//:go_default_library"],
)

go_binary(
    name = "proxy",
    embed = [":proxy_lib"],
    visibility = ["//visibility:public"],
)
