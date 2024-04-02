package handler

import (
	"fmt"
	"htmxx/service"
	"htmxx/templ"
	"net/http"
	"strconv"
)

type TimelineHandler struct {
	timelineService service.TimelineService
	userService     service.UserService
}

func (h *TimelineHandler) GetTimeline(w http.ResponseWriter, r *http.Request) {
var user = h.userService.GetCurrentUser(r)
	minid, iderr := strconv.Atoi(r.PathValue("minid"))
	if iderr != nil {
		minid = 0
	}
	timeline, err := h.timelineService.GetTimeline(user, minid)
	if err != nil {
		// handle error
		http.Error(w, "Error: %v", http.StatusInternalServerError)
		return
	}
	// handle timeline
	timelineComponent := templ.Timeline(timeline, user)
	rerr := templ.Layout(timelineComponent, "Home", true).Render(r.Context(), w)

	if rerr != nil {
		http.Error(w, rerr.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TimelineHandler) GetUserTimeline(w http.ResponseWriter, r *http.Request) {
	requster := h.userService.GetCurrentUser(r)
	author := r.PathValue("author")
	minid, err := strconv.Atoi(r.PathValue("minid"))
	if err != nil {
		minid = 0
	}

	timeline, err := h.timelineService.GetUserTimeline(requster, author, minid)
	if err != nil {
		// handle error
		http.Error(w, "Error: %v", http.StatusInternalServerError)
		return
	}
	if len(timeline) == 0 {
		// handle error
		http.Error(w, "User not found", http.StatusNoContent)
		return
	}

	timelineComponent := templ.UserTimeline(timeline, author, requster)
	rerr := templ.Layout(timelineComponent, fmt.Sprintf("%s's Timeline", author), true).Render(r.Context(), w)

	if rerr != nil {
		http.Error(w, rerr.Error(), http.StatusInternalServerError)
		return
	}
}
