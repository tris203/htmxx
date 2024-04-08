package main

import (
	"bytes"
	"fmt"
	"htmxx/model"
	"htmxx/service"
	"htmxx/templ"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (h *application) GetTweet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		// handle error
		http.Error(w, "Invalid tweet ID", http.StatusBadRequest)
		return
	}
	tweet, err := h.GetTweetData(id, r.Context())
	if err != nil {
		// handle error
		fmt.Println(err)
		http.Error(w, "Tweet not found", http.StatusNotFound)
		return
	}
	// handle tweet
	tweetComponent := templ.IndivTweet(tweet, r.Context().Value("user").(string))
	err = templ.Layout(tweetComponent, "Tweet", false).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *application) CreateTweet(w http.ResponseWriter, r *http.Request) {
	var content = r.FormValue("content")
	// envrypt the ip
	author := r.Context().Value("user").(string)

	tweet := &model.Tweet{
		Author:  author,
		Content: content,
	}

	newid, err := h.CreateTweetData(tweet, r.Context())
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tweet.ID = newid
	tweet.Created = time.Now().UTC()
	// handle success
	var stringTmpl bytes.Buffer
	insertTweetComponent := templ.InsertTweet(tweet)
	insertTweetComponent.Render(r.Context(), &stringTmpl)

	h.eventsService.AddMessage(service.Event{Data: fmt.Sprintf("%s", strings.ReplaceAll(stringTmpl.String(), "\n", "")), EventName: "new-tweet"})
	http.Error(w, "", http.StatusCreated)
	confirmationContainer := templ.CreateNewTweetConfirmation()
	confirmationContainer.Render(r.Context(), w)

}

func (h *application) AddLike(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newLikeCount, likedBySelf, err := h.AddLikeData(id, r.Context())
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// handle success
	h.eventsService.AddMessage(service.Event{Data: fmt.Sprintf("%d", newLikeCount), EventName: fmt.Sprintf("new-like-count-%d", id)})
	heartComponent := templ.Heart(id, likedBySelf, newLikeCount, true)
	// fmt.Fprintf(w, "%d", newLikeCount)
	heartComponent.Render(r.Context(), w)
}

func (h *application) RemoveLike(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newLikeCount, likedBySelf, err := h.RemoveLikeData(id, r.Context())
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// handle success
	h.eventsService.AddMessage(service.Event{Data: fmt.Sprintf("%d", newLikeCount), EventName: fmt.Sprintf("new-like-count-%d", id)})
	heartComponent := templ.Heart(id, likedBySelf, newLikeCount, false)
	heartComponent.Render(r.Context(), w)
}

func (h *application) DeleteTweet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	deletedID, err := h.DeleteTweetData(id, r.Context()) 
	if err != nil || deletedID == 0 {
		// handle error
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	// handle success
	deleteTweetComponent := templ.DeletedTweet(deletedID)
	deleteTweetComponent.Render(r.Context(), w)

}
