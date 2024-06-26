package main

import (
	"fmt"
	"htmxx/db"
	"htmxx/model"
	"htmxx/templ"
	"net/http"
)

func shapeDBTweets(tweets []db.Tweet) []*model.Tweet {
	// TODO: details around bools, although this is a search handler so it might be fine here due to compct rendering
	var shapedTweets []*model.Tweet
	for _, tweet := range tweets {
		shapedTweets = append(shapedTweets, &model.Tweet{
			ID:               tweet.TweetID,
			Author:           tweet.Author,
			Content:          tweet.Content,
			Created:          tweet.Created.Time,
		})
	}
	return shapedTweets
}

func (h *application) Search(w http.ResponseWriter, r *http.Request) {

	searchTerm := r.FormValue("search")

	if len(searchTerm) == 0 {
		resetComponent := templ.ResetSearch()
		resetComponent.Render(r.Context(), w)
		return
	}


	searchResults, err := h.query.SearchTweets(r.Context(), fmt.Sprintf("%%%s%%", searchTerm))

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

func (h *application) GetSearch(w http.ResponseWriter, r *http.Request) {

	rerr := templ.Layout(templ.SearchForm(), "Search", true).Render(r.Context(), w)

	if rerr != nil {
		http.Error(w, rerr.Error(), http.StatusInternalServerError)
	}
}
