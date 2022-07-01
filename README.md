# Go Vite Kit

Minimal boilerplate project for a Golang server using [Fiber](https://github.com/gofiber/fiber) and [ViteJS](https://vitejs.dev/) for static pages

## Features

-   ⚡️ [Go Fiber](https://github.com/gofiber/fiber)
-   📦 [ViteJS](http://vitejs.dev/)
-   🎨 [Sass](https://sass-lang.com/)
-   🗄️ [Sqlite3](https://github.com/mattn/go-sqlite3)

## Architecture

-   `_frontend/`

    This is a Vite project for building all the static pages used by this app.

    The `routes.js` (this is used both from `server.js` for _serving_ and from `vite.config.js` for _building_) file contains a mapping from express route patterns to entry html files, this is useful for rendering the same page for multiple urls in development mode.

-   `database/`

    Module with a `Database` interface and two implementation: `memDB` is an in-memory database for testing purposes. `sqliteDB` is a wrapper for working with an SQLite database.

-   `routes/`

    Various functions for configuring all the server routes.

## Usage

First install the required npm packages

```bash
$ cd _frontend
_frontend/ $ npm install
```

### Development

```bash
# Development
$ MODE=dev go run -v .

# Development with watcher
$ fd -e go | MODE=dev entr -r go run -v .
```

### Production

First build the `_frontend/dist` folder using

```bash
$ cd _frontend
$ npm run build
```

and then

```bash
# Production
$ go run -v .
```
