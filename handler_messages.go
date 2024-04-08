package main

import (
	"htmxx/templ"
	"net/http"
)

func (h *application) GetMessages(w http.ResponseWriter, r *http.Request) {
	// handle messages
	messagesComponent := templ.Messages()
	err := templ.Layout(messagesComponent, "Messages", false).Render(r.Context(), w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

