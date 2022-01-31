# Notebook web service

## Development

To run the service with edit-refresh for the UI, you need to run the following
combination of services:

1. [UI service](ui)
1. [API service](server)
1. [reverse HTTP proxy](proxy)

Run each service in a separate terminal.

### Run the UI service

```sh
$ cd ui
$ yarn install
$ yarn run dev
```

This runs `preact watch` which defaults to `localhost:8080`.

### Run the API service

```sh
$ cd server
$ go build .
$ ./server -port 5000
```

> Note: if you are seeing errors because port 5000 is in use, or you are getting
> permission errors when accessing the server on this port, you should be aware
> that port 5000 is used by Control Center in macOS Monterey for AirPlay.
>
> You can either use a different port or [disable
> AirPlay](https://nono.ma/port-5000-used-by-control-center-in-macos) by
> unchecking it in "System Preferences › Sharing" if you don't use it.

### Run the proxy

Download and build the proxy:

```
$ cd proxy
$ go build .
```

Use the provided [`config.yaml`](proxy/config.yaml) which will use the above
default ports, or update `config.yaml` locally to suit your needs:

```
$ tail -5 config.yaml
routes:
  - path: /api/
    target: http://localhost:5000
  - path: /
    target: http://localhost:8080
```

Run the proxy:

```
$ ./proxy -config config.yaml -port 9000
```

Now you can access http://localhost:9000 to access the combination of servers
through the reverse proxy.

## Production

For emulating production, we can simply pre-build the UI for deployment, and
then point our Go API server to the pre-built UI output, and thus, we won't need
to run the Node server in production, which simplifies our production
deployment.

### Build the web UI

```
$ cd ui
$ yarn run build
```

### Run the API service

```
$ cd server
$ go build .
$ ./server -web-root ../ui/build -port 5000
```

Now you can access http://localhost:5000 which will serve both UI and API
requests.
