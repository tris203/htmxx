package handler

import (
	"htmxx/templ"
	"htmxx/service"
	"net/http"
)

type ProfileHandler struct {
	userService service.UserService
}

func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	curUser := h.userService.GetCurrentUser(r)
	profileComponent := templ.Profile(curUser)
	templ.Layout(profileComponent, "Profile", false).Render(r.Context(), w)
}
