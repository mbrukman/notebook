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

### Run the proxy

Download and build the proxy:

```
$ go get -u github.com/mbrukman/notebook/web/proxy
```

Create config matching the port numbers used above:

```
$ cat > config.yaml <<EOF
routes:
  - path: /api/
    target: http://localhost:5000
  - path: /
    target: http://localhost:8080
EOF
```

Run the proxy with the config we created:

```
$ proxy -config config.yml -port 9000
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
