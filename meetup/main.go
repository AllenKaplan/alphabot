package meetup

import (
	"context"
	"errors"
	"github.com/AllenKaplan/alphabot/meetup/handler"
	proto "github.com/AllenKaplan/alphabot/meetup/proto"
	"strings"
	"time"
)

type MeetupService struct {
	context.Context
	repo meetup.MeetupRepo
}

func NewMeetupClient(ctx context.Context) MeetupService {
	return MeetupService{
		Context: ctx,
		repo:    meetup.NewRepo(),
	}
}

func (srv MeetupService) CreateMeetup(req string) (string, error) {
	params := strings.Fields(req)

	meetupTime, err := time.Parse("2006-01-02", params[2])
	if err != nil {
		return "", err
	}

	meetupToInsert := &proto.Meetup{
		Name:     params[0],
		Location: params[1],
		Time:     meetupTime,
	}

	created, err := srv.repo.CreateMeetup(*meetupToInsert)
	if err != nil {
		return "", err
	}

	if !created {
		return "", errors.New("go.bot.meetup.CreateMeetup | Could not create meetup")
	}

	return "Meetup Created", nil
}

func (srv MeetupService) GetMeetup(req string) (*proto.Meetup, error) {
	res, err := srv.repo.GetMeetup(req)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errors.New("No meetup found by req: " + req)
	}

	return res, nil
}
