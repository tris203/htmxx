package service

import (
	"htmxx/dummy"
	"htmxx/model"
)

type TimelineService struct {
}

func (s *TimelineService) GetTimeline(userid string, minid int) ([]*model.Tweet, error) {
	var allTimelines, err = dummy.GetTimeline(userid, minid)
	if err != nil {
		return nil, err
	}
	return allTimelines, nil
}

func (s *TimelineService) GetUserTimeline(userid string, author string, minid int) ([]*model.Tweet, error) {
	var userTimelines, err = dummy.GetUserTimeline(userid, author, minid)
	if err != nil {
		return nil, err
	}
	return userTimelines, nil
}
