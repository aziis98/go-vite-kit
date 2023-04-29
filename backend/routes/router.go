package routes

import "go-vite-kit/backend/database"

type Router struct {
	Database database.Database
}

type htmlEntrypoint struct {
	Route    string `json:"route"`
	Filename string `json:"filename"`
}

var HtmlEntrypoints = []htmlEntrypoint{
	{"/", "./index.html"},
	// {"/u/:username", "./user.html"},
}
