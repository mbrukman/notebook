# File viewer

This is a simple web server which renders text files (including HTML, JS, CSS)
and Markdown files (dynamically rendered to HTML). Some use cases include
quickly browsing local directories via a web browser, including Markdown files
with relative links.

This is an early prototype, but you're welcome to play around with it and
provide feedback, potential new use cases, etc.

## Building

Build in the current directory:

```sh
$ go build .
```

Or you can build it from the top of the tree in a local checkout:

```sh
$ go build ./web/filevewer
```

Or you can build via Bazel:

```sh
$ bazel build //web/fileviewer
```

Or you can build it without having this repo locally:

```sh
$ go install github.com/mbrukman/notebook/web/fileviewer@latest
```

# Running

Run locally with a custom web root (only accessible from `localhost` by
default):

```sh
$ ./fileviewer -web-root ~/notebook
```

Expose it to everyone who can access this computer via the network:

```sh
$ ./fileviewer -web-root ~/notebook -host 0.0.0.0
```

Get a list of available flags:

```sh
$ ./fileviewer -help
```

Running via Bazel (this path is printed when you run the `bazel build ...`
command above):

```sh
$ bazel-bin/web/fileviewer/fileviewer_/fileviewer [...flags]
```

