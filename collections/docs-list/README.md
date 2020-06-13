# Docs list

Open `static/index.html` in your browser to see the sortable/filterable docs
list prototype with placeholder data. See below for generating your own data set
for this list.

A dynamic server version of the same list is in-progress in the
[`dynamic`](dynamic) directory.

## Using with your own data

Here's how to generate a custom version of `data.js` from your own files:

1. Visit
   [Google Drive API](https://developers.google.com/drive/api/v3/reference/files/list)
   reference

1. Open the "Try this API" side panel

1. in the `q` textbox, fill in a query such as:

   ```
   '<your-email>' in owners AND createdTime >= '2020-01-01'
   ```

1. in the "fields" textbox, fill in a filter such as:

   ```
   files(name, createdTime, webViewLink)
   ```

1. click on the button labeled "Execute" at the bottom of the page, and
   authorize it via OAuth if needed

1. make sure your query returns HTTP 200; if not, fix the query or
   authentication and try again

1. click inside the text box with the output from the query

1. select everything (Ctrl-A / ⌘-A) in that text box

1. copy (Ctrl-C / ⌘-C) the selection

1. update the file `docs-list/data.js` as follows:

   ```
   var DOCS = {...paste (Ctrl-V / ⌘-V) output from the query here...};
   ```

Now you can open `static/index.html` in your browser and interact with the list!
