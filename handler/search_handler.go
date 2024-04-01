package handler

import (
	"htmxx/dummy"
	"htmxx/templ"
	"net"
	"net/http"
)

type SearchHandler struct {
}

func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	// handle about
	searchTerm := r.FormValue("search")

	if len(searchTerm) == 0 {
		resetComponent := templ.ResetSearch()
		resetComponent.Render(r.Context(), w)
		return
	}

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	searchResults, err := dummy.SearchTweets(searchTerm, ip)

	if len(searchResults) == 0 {
		noResultsComponent := templ.NoResults(searchTerm)
		noResultsComponent.Render(r.Context(), w)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	searchResultsComponent := templ.SearchResult(searchTerm, searchResults)
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
