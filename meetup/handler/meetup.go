package meetup

import (
	"errors"

	"github.com/AllenKaplan/alphabot/meetup/proto"
)

type MeetupRepo struct{}

func NewMeetupRepo() *MeetupRepo {
	return &MeetupRepo{}
}

var (
	meetups []*meetup.Meetup
)

type Repo interface {
	CreateMeetup(*meetup.Meetup) (bool, error)
	GetMeetup(string) (*meetup.Meetup, error)
	GetAllMeetups() ([]*meetup.Meetup, error)
}

func (repo MeetupRepo) CreateMeetup(meetup *meetup.Meetup) (bool, error) {
	meetups = append(meetups, meetup)
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
