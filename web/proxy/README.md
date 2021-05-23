# Proxy

This is a simple [reverse HTTP proxy][reverse-proxy], which you can run in front
of several servers to enable routing to backends based on URL patterns. The
primary use case it enables is running our API server (Go) and UI server (Node)
behind a single URL, so that APIs are easily accessible from the UI as they are
served from the same origin.

This lets us easily mimic our (expected) production environment during
development.

## Install

If you're using Go 1.11 or 1.12, you need to set this environment variable
first:

```sh
$ export GO111MODULE=on
```

Then, install the proxy as follows:

```sh
$ go get -u github.com/mbrukman/notebook/web/proxy
```

This will install the binary in `$GOPATH/bin`; `$GOPATH` defaults to `$HOME/go`
if not set.

## Run

Assuming `$GOPATH/bin` is in your `$PATH`, you can run it as follows:

```sh
$ proxy -config config.yml
```

Run `proxy -help` to see the complete list of available flags; in particular,
you may want to set one or more of the following:

* `-host [host]` – empty string or `0.0.0.0` will make the proxy accessible from
  other hosts; `localhost` or `127.0.0.1` (default) will make the proxy only
  accessible from the same machine as the proxy, for security reasons. The proxy
  will print at start time the host it's listening on.

* `-port [number]` – change the port on which the proxy will listen; `9000` is
  the default. The proxy will print at start time which port it's listening
  on.

[reverse-proxy]: https://en.wikipedia.org/wiki/Reverse_proxy
