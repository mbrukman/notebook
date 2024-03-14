# Notebook

[![Bazel][bazel-ci-badge]][bazel-ci-url]
[![Node][node-ci-badge]][node-ci-url]
[![Go][go-ci-badge]][go-ci-url]
[![Go docs][go-doc-badge]][go-doc-url]
[![Go Report Card][go-report-card-badge]][go-report-card-url]

[bazel-ci-badge]: https://github.com/mbrukman/notebook/actions/workflows/bazel.yaml/badge.svg
[bazel-ci-url]: https://github.com/mbrukman/notebook/actions/workflows/bazel.yaml
[node-ci-badge]: https://github.com/mbrukman/notebook/actions/workflows/node.yaml/badge.svg
[node-ci-url]: https://github.com/mbrukman/notebook/actions/workflows/node.yaml
[go-ci-badge]: https://github.com/mbrukman/notebook/actions/workflows/go.yaml/badge.svg
[go-ci-url]: https://github.com/mbrukman/notebook/actions/workflows/go.yaml
[go-doc-badge]: http://img.shields.io/badge/godoc-reference-informational.svg
[go-doc-url]: https://pkg.go.dev/github.com/mbrukman/notebook
[go-report-card-badge]: https://goreportcard.com/badge/github.com/mbrukman/notebook
[go-report-card-url]: https://goreportcard.com/report/github.com/mbrukman/notebook

## Overview

This digital notebook is for keeping track of semi-structured data, such as
lists and collections of URLs, for sharing and reference.

This is a work-in-progress; here are some ongoing efforts:

* Preact-based [web UI](web/ui/#readme) and [Go API server](web/server)

  You can run the web UI [directly](web/ui/#readme), or you can use the
  provided [Dockerfile](docker) to build and run the container locally.

* [collections UI](prototypes/collections/#readme) prototype

## Contributing

See [`CONTRIBUTING.md`](CONTRIBUTING.md) for details.

## License

Apache 2.0; see [`LICENSE`](LICENSE) for details.

## Disclaimer

This project is not an official Google project. It is not supported by Google
and Google specifically disclaims all warranties as to its quality,
merchantability, or fitness for a particular purpose.
