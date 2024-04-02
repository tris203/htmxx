package main

import (
	"fmt"
	"htmxx/handler"
	"net/http"
	"os"
)

func main() {
	router := http.NewServeMux()
	tweetHandler := handler.TweetHandler{}
	timelineHandler := handler.TimelineHandler{}
	sseEventsHandler := handler.SSEEventsHandler{}
	aboutHandler := handler.AboutHandler{}
	searchHandler := handler.SearchHandler{}
	messageHandler := handler.MessagesHandler{}
	likesListHandler := handler.LikesListHandler{}
	bookmarkHandler := handler.BookmarkHandler{}
	profileHandler := handler.ProfileHandler{}
	staticFS := http.FileServer(http.Dir("./static"))
	router.HandleFunc("GET /", timelineHandler.GetTimeline)
	router.HandleFunc("GET /{author}/", timelineHandler.GetUserTimeline)
	router.HandleFunc("GET /page/{author}/{minid}/", timelineHandler.GetUserTimeline)
	router.HandleFunc("GET /page/{minid}", timelineHandler.GetTimeline)
	router.HandleFunc("GET /tweet/{id}", tweetHandler.GetTweet)
	router.HandleFunc("DELETE /tweet/{id}", tweetHandler.DeleteTweet)
	router.HandleFunc("POST /like/{id}", tweetHandler.AddLike)
	router.HandleFunc("POST /tweet", tweetHandler.CreateTweet)
	router.Handle("GET /static/", http.StripPrefix("/static/", staticFS))
	router.HandleFunc("GET /events", sseEventsHandler.EventsHandler)
	router.HandleFunc("GET /about", aboutHandler.GetAbout)
	router.HandleFunc("GET /search", searchHandler.GetSearch)
	router.HandleFunc("POST /search", searchHandler.Search)
	router.HandleFunc("GET /messages/", messageHandler.GetMessages)
	router.HandleFunc("GET /likes/", likesListHandler.GetLikesList)
	router.HandleFunc("GET /likes/{author}/", likesListHandler.GetLikesList)
	router.HandleFunc("GET /bookmarks/", bookmarkHandler.GetBookmark)
	router.HandleFunc("POST /bookmark/{id}", bookmarkHandler.AddBookmark)
	router.HandleFunc("GET /profile/", profileHandler.GetProfile)



	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Printf("Listening on port %s\n", port)
	server.ListenAndServe()
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	//redirect to timeline
	http.Redirect(w, r, "/timeline", http.StatusSeeOther)
}
