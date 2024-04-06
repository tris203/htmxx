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
}

func (h *TimelineHandler) GetTimeline(w http.ResponseWriter, r *http.Request) {
var user = r.Context().Value("user").(string)
	minid, iderr := strconv.ParseInt(r.PathValue("minid"), 10, 64)
	if iderr != nil {
		minid = 9223372036854775807
	}
	timeline, err := h.timelineService.GetTimeline(minid, user)
	if err != nil {
		// handle error
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
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
	requster := r.Context().Value("user").(string)
	author := r.PathValue("author")
	minid, err := strconv.ParseInt(r.PathValue("minid"), 10, 64)
	if err != nil {
		minid = 9223372036854775807
	}

	timeline, err := h.timelineService.GetUserTimeline(author, minid, requster)
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
