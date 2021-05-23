# API server

## Run API service only

Running just the API server without any static data (HTML, JS, CSS):

```
$ go build .
$ ./server -port 5000
```

You will also need to run the UI server (Node) as well as the proxy; see
[`../README.md`](../README.md) for complete details on how to do this. This
approach is best for local development, as the Node server provides edit-refresh
capability with automatic live rebuilds for any changes to static data.

## Run complete service

Run the API server with complete build of static data (HTML, JS, CSS):

```
$ cd ../ui
$ yarn install
$ yarn run build

$ cd ../server
$ go build .
$ ./server -web-root ../ui/build -port 5000
```

You can now access http://localhost:5000 without a proxy or Node server for the
complete application. However, this does not provide the edit-refresh
capability, so this is good for production deployment were the HTML, CS, and CSS
resources will not be updated at runtime.
