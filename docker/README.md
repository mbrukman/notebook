# Running Notebook via Docker

> Note: if you run into any issues with the commands below and need to debug the
> installation, the container, or want to run some custom scripts, you can easily
> get a shell inside the container by running:
>
> ```
> $ ./docker.sh shell
> ```

## Status

The current implementation supports creating notes (consisting of a title and
body), but does not persist the data anywhere, not even in-memory of the
server: it's only stored client-side, in the local memory state of your browser
tab, so even opening another tab will reset your state to a clean slate.

## Instructions

1. Build the container

   ```
   $ ./docker.sh build
   ```

2. Install the required Node dependencies

   ```
   $ ./docker.sh install
   ```

3. Run the container (builds the app and runs the server)

   ```
   $ ./docker.sh run
   ```

4. Access the Notebook UI

   Visit http://localhost:8080/ to see the UI.

