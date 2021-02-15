# Notebook frontend

To run the frontend:

```
$ yarn install
$ yarn run dev
```

and open the URL printed on the screen.

## Installation notes

On Ubuntu 20.04 LTS, the `/usr/bin/node` binary provided by the `nodejs`
package is version `10.19.0`, which unfortunately does not work well with
Preact CLI 3.0.5; you will get the following error:

```
TypeError: options.map(...).flat is not a function
```

when attempting to run this app. This is documented in
https://github.com/preactjs/preact-cli/issues/1500 and fixed in
https://github.com/preactjs/preact-cli/pull/1493, but this was merged after the
most-recent (at the time of this writing) release of Preact CLI 3.0.5, and
hence, there are no later releases that fix this issue.

One way to address this issue (while waiting for a later version of Preact CLI
to be released) is to install a newer version of Node, which we can do via
[Ubuntu Snap](https://snapcraft.io/node):

```
$ sudo snap install node --channel=15/stable --classic
```
