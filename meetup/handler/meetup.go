package meetup

import (
	"errors"

	"github.com/AllenKaplan/alphabot/meetup/proto"
)

type MeetupRepo struct{}

var (
	meetups []*meetup.Meetup
)

func NewRepo() MeetupRepo {
	return MeetupRepo{}
}

func (repo MeetupRepo) CreateMeetup(meetup meetup.Meetup) (bool, error) {
	meetups = append(meetups, &meetup)
	return true, nil
}

func (repo MeetupRepo) GetMeetup(name string) (*meetup.Meetup, error) {
	for _, currentMeetup := range meetups {
		if name == currentMeetup.Name {
			return currentMeetup, nil
		}
	}

	return nil, errors.New("go.bot.meetup.repo.GetMeetup - Meetup not found")
}

func (repo MeetupRepo) GetAllMeetups() ([]*meetup.Meetup, error) {
	return meetups, nil
}
