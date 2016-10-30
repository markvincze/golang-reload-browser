# Golang page reloader

This is a small sample application implemented in Golang which can programmatically refresh a tab in a browser.
It is accompanying a blog post I've written about the subject with all the details: .

It hosts a small WebSocket service with a single `reload` endpoint, to which we can connect from the browser, and send a message every time we want it to be reloaded.

## Usage

First, compile the application by running the script `build.sh`, it'll put the output in the file `bin/reload`.

```
./build.sh
```

Then start up the application.

```
bin/reload
```

Then open the static file `reloadtest.html` in the browser.

Every time we press Enter in the terminal, the page should display a message. If we uncomment the `location.reload()` call, then instead of printing a message, the tab will be refreshed.
