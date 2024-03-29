package handler

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"htmxx/templ"
	"htmxx/model"
	"htmxx/service"
	"net"
	"net/http"
	"strconv"
	"strings"
)

type TweetHandler struct {
	tweetService  service.TweetService
	eventsService service.EventsService
}

func (h *TweetHandler) GetTweet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		// handle error
		http.Error(w, "Invalid tweet ID", http.StatusBadRequest)
		return
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	tweet, err := h.tweetService.GetTweet(id, ip)
	if err != nil {
		// handle error
		http.Error(w, "Tweet not found", http.StatusNotFound)
		return
	}
	// handle tweet
	tweetComponent := templ.Tweet(tweet)
	err = templ.Layout(tweetComponent, "Tweet", false).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *TweetHandler) CreateTweet(w http.ResponseWriter, r *http.Request) {
	var content = r.FormValue("content")
	var ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	// envrypt the ip
	author := fmt.Sprintf("%x", md5.Sum([]byte(ip)))

	tweet := &model.Tweet{
		Author:  author,
		Content: content,
	}

	_, err := h.tweetService.CreateTweet(tweet)
	// show delete on first render
	tweet.IsAuthor = true
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// handle success
	var stringTmpl bytes.Buffer
	insertTweetComponent := templ.InsertTweet(tweet)
	insertTweetComponent.Render(r.Context(), &stringTmpl)

	h.eventsService.AddMessage(service.Event{Data: fmt.Sprintf("%s", strings.ReplaceAll(stringTmpl.String(), "\n", "")), EventName: "new-tweet"})
	http.Error(w, "<span id='tweet-success' class='text-green-700' remove-me='3s' hx-ext='remove-me'>Tweet created</span>", http.StatusCreated)

}

func (h *TweetHandler) AddLike(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newLikeCount, likedBySelf, err := h.tweetService.AddLike(id, ip)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// handle success
	h.eventsService.AddMessage(service.Event{Data: fmt.Sprintf("%d", newLikeCount), EventName: fmt.Sprintf("like-count-%d", id)})
	svgColor := ""
	if likedBySelf {
		svgColor = "fill-red-500"
	} else {
		svgColor = ""
	}
	fmt.Fprintf(w, "%d", newLikeCount)
	fmt.Fprintf(w, `
  <svg
    id="like-icon-%d"
	hx-swap-oob="true"
    xmlns="http://www.w3.org/2000/svg"
    width="24"
    height="24"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    class="w-4 h-4 %s"
  >
    <path
      d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"
    ></path>
  </svg>
`, id, svgColor)

}

func (h *TweetHandler) DeleteTweet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	tweet, err := h.tweetService.DeleteTweet(id, ip)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// handle success
	deleteTweetComponent := templ.DeletedTweet(tweet)
	deleteTweetComponent.Render(r.Context(), w)

}
