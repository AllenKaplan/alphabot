package meetup

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type MeetupService struct{
	context.Context
}

type Meetup struct {
	Name string
	Location string
	Time time.Time
}

func (m Meetup) String() string {
	return fmt.Sprintf("%s | %s | %s", m.Name, m.Location, m.Time)
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
		Name:     params[0],
		Location: params[1],
		Time:     meetupTime,
	}

	res, err := Cr
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
