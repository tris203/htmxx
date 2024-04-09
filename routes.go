package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	staticFS := http.FileServer(http.Dir("./static"))
	mux.HandleFunc("GET /{author}/", app.GetUserTimeline)
	mux.Handle("GET /static/", http.StripPrefix("/static", staticFS))
	mux.HandleFunc("GET /events", app.EventsHandler)
	mux.HandleFunc("GET /", app.GetTimeline)
	mux.HandleFunc("GET /about", app.GetAbout)
	mux.HandleFunc("GET /messages", app.GetMessages)
	mux.HandleFunc("GET /profile/", app.GetProfile)
	mux.HandleFunc("POST /tweet", app.CreateTweet)
	mux.HandleFunc("GET /tweet/{id}/", app.GetTweet)
	mux.HandleFunc("DELETE /tweet/{id}", app.DeleteTweet)
	mux.HandleFunc("POST /like/{id}/", app.AddLike)
	mux.HandleFunc("DELETE /like/{id}/", app.RemoveLike)
	mux.HandleFunc("GET /search", app.GetSearch)
	mux.HandleFunc("POST /search", app.Search)
	mux.HandleFunc("GET /likes/", app.GetLikesList)
	mux.HandleFunc("GET /likes/{author}/", app.GetLikesList)
	mux.HandleFunc("GET /bookmarks/", app.GetBookmark)
	mux.HandleFunc("POST /bookmark/{id}/", app.AddBookmark)
	mux.HandleFunc("DELETE /bookmark/{id}/", app.RemoveBookmark)
	mux.HandleFunc("GET /page/{minid}/", app.GetTimeline)
	mux.HandleFunc("GET /page/{author}/{minid}/", app.GetUserTimeline)

	return mux
}
