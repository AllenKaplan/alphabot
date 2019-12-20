package meetup

import (
	"fmt"
	"github.com/AllenKaplan/alphabot/meetup"
)

type MeetupRepo struct {}

var (
	meetups []meetup.Meetup
)

type MeetupService interface {
	CreateMeetup(meetup.Meetup) string
	GetMeetup(string) meetup.Meetup
}

func (repo MeetupRepo) CreateMeetup(meetup meetup.Meetup) string {
	meetups = append(meetups, meetup)
	return fmt.Sprintf("Added meetup: %v", meetup)
}

func (repo MeetupRepo)  GetMeetup(name string) string {
	for _, currentMeetup := range meetups {
		if name == currentMeetup.Name {
			return currentMeetup
		}
	}

	return "Could not find meetup"
}

