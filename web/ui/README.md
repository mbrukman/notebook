# Notebook web UI

To run the web UI:

```sh
$ yarn install
$ yarn run dev
```

and open the URL printed on the screen.

Note that the web UI makes API calls to an [API server](../server), and since
the API server will run on a different port, both the UI and API server need to
run behind a [proxy](../proxy).

You can either run the whole combination to match what it would look like in
production, or you can run with a fake API implementation:

```sh
$ FAKE_API=1 yarn run dev
```

or:

```sh
$ make dev FAKE_API=1
```

## Installation notes

On Ubuntu 20.04 LTS, the `/usr/bin/node` binary provided by the `nodejs` package
is version 10.19.0, which unfortunately does not work well with Preact CLI
3.0.5. When attempting to run this app with those package versions, you will see
the following error:

```
TypeError: options.map(...).flat is not a function
```

This is [documented][preact-cli-issue] and [fixed][preact-cli-fix], but it was
merged after the most-recent (at the time of this writing) release of Preact CLI
3.0.5 (9 Dec 2020), and hence, there are no later releases that fix this issue.

One way to address this issue (while waiting for a later version of Preact CLI
to be released) is to install a newer version of Node, which we can do via
[Ubuntu Snap][ubuntu-snap]:

```sh
$ sudo snap install node --channel=15/stable --classic
```

Snap will install the new `node` binary in `/snap/bin`, so make sure that this
directory appears before `/usr/bin` in your `$PATH` for the correct `node`
binary to be automatically selected. Additionally, you can uninstall the older
version of Node provided by the `nodejs` package to remove it from `/usr/bin`
entirely.

[preact-cli-issue]: https://github.com/preactjs/preact-cli/issues/1500
[preact-cli-fix]: https://github.com/preactjs/preact-cli/pull/1493
[ubuntu-snap]: https://snapcraft.io/node
