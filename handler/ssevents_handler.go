package handler

import (
	"fmt"
	"htmxx/service"
	// "math/rand"
	"net/http"
	// "time"
)

type SSEEventsHandler struct {
	eventsService service.EventsService
}

func (h *SSEEventsHandler) EventsHandler(w http.ResponseWriter, r *http.Request) {
	// handle SSE events
	fmt.Println("Client Connected")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	var channel = make(chan service.Event)
	h.eventsService.AddChannel(&channel)
	defer func() {
		close(channel)
	}()

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
	}

	for {
		select {
		case message := <-channel:
			_, err := fmt.Fprintf(w, "event: %s\ndata: %s\n\n", message.EventName, message.Data)
			if err != nil {

				fmt.Println(fmt.Printf("Error sending message: %v", err))
			}
			flusher.Flush()
		case <-r.Context().Done():
			h.eventsService.RemoveChannel(&channel)
			return
		}
	}




	/* // go func() {
		for data := range channel {
			fmt.Println(fmt.Sprintf("Sending message: %s for %s", data.Data, data.EventName))
			_, err := fmt.Fprintf(w, "event: %s\ndata: %s\n\n", data.EventName, data.Data)
			if err != nil {
				fmt.Println(fmt.Printf("Error sending message: %v", err))
			}
			w.(http.Flusher).Flush()
		} */
	// }()

	/* // Simulate sending messages
	i := 10
	for {
		randomTweet := rand.Intn(3) + 1
		i++
		fmt.Fprintf(w, "event: like-count-%d\ndata: %d\n\n", randomTweet, i)
		time.Sleep(2 * time.Second)
		w.(http.Flusher).Flush()
	} */

}
