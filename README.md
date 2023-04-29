# Go Vite Kit

Minimal boilerplate project for a Golang server using [Fiber](https://github.com/gofiber/fiber) and [ViteJS](https://github.com/vitejs/vite) for static pages

## Features

-   ⚡️ [Go Fiber](https://github.com/gofiber/fiber)
-   📦 [ViteJS](http://vitejs.dev/)
-   🎨 [Sass](https://sass-lang.com/)
-   🗄️ [Sqlite3](https://github.com/mattn/go-sqlite3)

## Architecture

-   `frontend/`

    This is a Vite project for building all the static pages used by this app.

-   `backend/`

    This keeps all server related files

    -   `config/`

        Loads env variables and keeps them as globals 

    -   `database/`

        Module with a `Database` interface and two implementation: `memDB` is an in-memory database for testing purposes. `sqliteDB` is a wrapper for working with an SQLite database.

    -   `routes/`

        Various functions for configuring all the server routes.

        A very important file is `backend/routes/router.go` that contains the `HtmlEntrypoints` variable that is used both by the backend and ViteJS to mount HTML entrypoints. 
        
        When building the frontend ViteJS will call `go run ./cmd/routes` to read the content of this variabile, while during development a special route called `/dev/routes` gets mounted on the backend server and this lets Vite add all necessary entrypoints to the dev server.

## Usage

First install the required npm packages

```bash
$ npm install
```

### Development

```bash
# Development
$ MODE=dev go run -v ./cmd/server

# Development with watcher
$ fd -e go | MODE=dev entr -r go run -v ./cmd/server
```

### Production

You can build everything with the following command, it will build first the frontend and then the backend and generate `./out/server`. 

```bash
# Build
$ go run -v ./cmd/build

# Run
$ ./out/server
```
