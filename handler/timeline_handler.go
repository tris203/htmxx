package handler

import (
	"html/template"
	"htmxx/service"
	"net"
	"net/http"
	"strconv"
)

type TimelineHandler struct {
	timelineService service.TimelineService
}

func (h *TimelineHandler) GetTimeline(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	minid, err := strconv.Atoi(r.PathValue("minid"))
	if err != nil {
		minid = 0
	}
	timeline, err := h.timelineService.GetTimeline(ip, minid)
	if err != nil {
		// handle error
		http.Error(w, "Error: %v", http.StatusInternalServerError)
		return
	}
	// handle timeline
	tmpl, err := template.ParseFiles("templates/base.html", "templates/index.html", "templates/createTweet.html", "templates/timeline.html", "templates/tweet.html", "templates/likeButton.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, timeline)
}

func (h *TimelineHandler) GetUserTimeline(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	author := r.PathValue("author")
	minid, err := strconv.Atoi(r.PathValue("minid"))
	if err != nil {
		minid = 0
	}

	timeline, err := h.timelineService.GetUserTimeline(ip, author, minid)
	if err != nil {
		// handle error
		http.Error(w, "Error: %v", http.StatusInternalServerError)
		return
	}
	if len(timeline) == 0 {
		// handle error
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// handle timeline
	tmpl, err := template.ParseFiles("templates/base.html", "templates/usertimeline.html", "templates/wrapusertimeline.html", "templates/tweet.html", "templates/likeButton.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, timeline)
}
