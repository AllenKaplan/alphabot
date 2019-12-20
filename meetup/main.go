package meetup

import (
	"context"
	"strings"
	"time"
)

type MeetupService struct{
	context.Context
}

type Meetup struct {
	name string
	location string
	time time.Time
}

func NewMeetupClient(ctx context.Context) MeetupService {
	return MeetupService{
		Context: ctx,
	}
}


func (s MeetupService) CreateMeetup(req string) (string, error) {
	params := strings.Split(req, " ")

	meetupTime, err := time.Parse("2006-01-01" ,params[2])
	if err != nil {
		return "", err
	}

	meetupToInsert := &Meetup{
		name:     params[0],
		location: params[1],
		time:     meetupTime,
	}

	res, err := s.CreateMeetup(meetupToInsert)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s MeetupService) GetMeetup(req string) (interface{}, interface{}) {
	res, err := s.GetMeetup(req)
	if err != nil {
		return "", err
	}

	return res, nil
}
