package handler

import (
	"github.com/AllenKaplan/alphabot/meetup/proto"
)

type MeetupService struct{
	repo MeetupRepo
}

func NewMeetupService() *MeetupService {
	return &MeetupService{}
}
type MeetupRepo interface {
	CreateMeetup(*meetup.Meetup) (bool, error)
	GetMeetup(string) (*meetup.Meetup, error)
	GetAllMeetups() ([]*meetup.Meetup, error)
}

func (s MeetupService) CreateMeetup(req *meetup.Meetup) (bool, error) {
	created, err := s.CreateMeetup(req)
	if err != nil {
		return false, err
	}
	return created, nil
}

func (s MeetupService) GetMeetup(name string) (*meetup.Meetup, error) {
	res, err := s.GetMeetup(name)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s MeetupService) GetAllMeetups() ([]*meetup.Meetup, error) {
	res, err := s.GetAllMeetups()
	if err != nil {
		return nil, err
	}

	return res, nil
}
