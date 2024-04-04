package handler

import (
	"fmt"
	"htmxx/db"
	"htmxx/model"
	"htmxx/service"
	"htmxx/templ"
	"net/http"
)

type SearchHandler struct {
	userService service.UserService
	dbService   service.DBService
}

func shapeDBTweets(tweets []db.Tweet) []*model.Tweet {
	// TODO: details around bools, although this is a search handler so it might be fine here due to compct rendering
	var shapedTweets []*model.Tweet
	for _, tweet := range tweets {
		shapedTweets = append(shapedTweets, &model.Tweet{
			ID:               int(tweet.TweetID),
			Author:           tweet.Author,
			Content:          tweet.Content,
			Created:          tweet.Created.Time,
		})
	}
	return shapedTweets
}

func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {

	searchTerm := r.FormValue("search")

	if len(searchTerm) == 0 {
		resetComponent := templ.ResetSearch()
		resetComponent.Render(r.Context(), w)
		return
	}

	ctx, queries, dbConn, dberr := h.dbService.Connect()
	if dberr != nil {
		return
	}
	defer dbConn.Close()

	searchResults, err := queries.SearchTweets(ctx, fmt.Sprintf("%%%s%%", searchTerm))

	if len(searchResults) == 0 {
		noResultsComponent := templ.NoResults(searchTerm)
		noResultsComponent.Render(r.Context(), w)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	searchResultsComponent := templ.SearchResult(searchTerm, shapeDBTweets(searchResults))
	searchResultsComponent.Render(r.Context(), w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *SearchHandler) GetSearch(w http.ResponseWriter, r *http.Request) {

	rerr := templ.Layout(templ.SearchForm(), "Search", true).Render(r.Context(), w)

	if rerr != nil {
		http.Error(w, rerr.Error(), http.StatusInternalServerError)
	}
}
