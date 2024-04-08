package main

import (
	"htmxx/templ"
	"net/http"
)

func (h *application) GetProfile(w http.ResponseWriter, r *http.Request) {
	curUser := r.Context().Value("user").(string)
	profileComponent := templ.Profile(curUser)
	templ.Layout(profileComponent, "Profile", false).Render(r.Context(), w)
}
