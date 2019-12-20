package meetup

import (
	"fmt"
	"github.com/AllenKaplan/alphabot/meetup"
)

var (
	meetups []meetup.Meetup
)

func CreateMeetup(meetup meetup.Meetup) string {
	meetups = append(meetups, meetup)
	return fmt.Sprintf("Added meetup: %v", meetup)
}

func GetMeetup(name string) string {
	for meetup := range meetups {
		if name == meetup.Name {
			return meetup
		}
	}

	return "Could not find meetup"
}