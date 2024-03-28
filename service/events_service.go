package service

import (
	"fmt"
)

type EventsService struct {
}

type Event struct {
	Data      string
	EventName string
}

var CHANNEL_STORE []*chan Event = make([]*chan Event, 0)

func (s *EventsService) RemoveChannel(ch *chan Event) {
    pos := -1
    storeLen := len(CHANNEL_STORE)
    for i, msgChan := range CHANNEL_STORE {
        if ch == msgChan {
            pos = i
        }
    }

    if pos == -1 {
        return
    }
    CHANNEL_STORE[pos] = CHANNEL_STORE[storeLen-1]
    CHANNEL_STORE = CHANNEL_STORE[:storeLen-1]
    fmt.Println("Connection remains: ", len(CHANNEL_STORE))
}

func (s *EventsService) AddChannel(channel *chan Event) {
	//add channel to store
	CHANNEL_STORE = append(CHANNEL_STORE, channel)
}

func (s *EventsService) AddMessage(message Event) {
	//add message to channel
	for _, channel := range CHANNEL_STORE {
		*channel <- message
	}
}

