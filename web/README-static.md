# Static Web Files

Files in the `static` directory will get embedded into the Flamenco Manager
executable, and served as static files via its web server.

- `make webapp-static` clears it out and builds the webapp there. It also ZIPs
  the Blender add-on, and places it in there as well.
- `make clean-webapp-static` just does the clearing of the files.

`static/emptyfile` exists just to make sure that `go:embed` inside `web_app.go`
has something to work with, even before any static files have been built.

# Running static flamenco.blender.org site locally

The [Flamenco website](https://flamenco.blender.org/) runs off of [Hugo](https://gohugo.io/).

To locally run the site, run `make devserver-website`. Then visit https://localhost:1313/ in a webbrowser.

Alternatively, [manually install Hugo](https://gohugo.io/getting-started/installing/). Then, from the `web/project-website` directory, start the server with:

```
hugo server -D
```
