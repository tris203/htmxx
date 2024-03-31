package handler

import (
	"htmxx/templ"
	"net/http"
)

type AboutHandler struct {
}

func (h *AboutHandler) GetAbout(w http.ResponseWriter, r *http.Request) {
	// handle about
	aboutComponent := templ.About()
	err := templ.Layout(aboutComponent, "About", false).Render(r.Context(), w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
