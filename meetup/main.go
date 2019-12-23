package meetup

import (
	"context"
	"errors"
	"github.com/AllenKaplan/alphabot/meetup/handler"
	proto "github.com/AllenKaplan/alphabot/meetup/proto"
	"regexp"
	"strings"
	"time"
)

type MeetupService struct {
	context.Context
	repo *meetup.MeetupRepo
}

func NewMeetupClient(ctx context.Context) MeetupService {
	return MeetupService{
		Context: ctx,
		repo:    meetup.NewMeetupRepo(),
	}
}

func (srv MeetupService) CreateMeetup(req string) (string, error) {
	meetupToInsert, err := parseMeetup(req)
	if err != nil {
		return "Error parsing meetup", err
	}

	created, err := srv.repo.CreateMeetup(meetupToInsert)
	if err != nil {
		return "", err
	}

	if !created {
		return "", errors.New("go.bot.meetup.CreateMeetup | Could not create meetup")
	}

	return "Meetup Created ", nil
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

func (srv MeetupService) GetAllMeetups(s string) ([]*proto.Meetup, error) {
	res, err := srv.repo.GetAllMeetups()
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errors.New("no meetups found")
	}

	return res, nil
}

func parseMeetup(s string) (*proto.Meetup, error) {
	regex := regexp.MustCompile("[^\\s\"][\\w\\s'-:]+")
	params := regex.FindAllString(s, -1)

	if len(params) < 2 {
		params = strings.Fields(s)

		if len(params) < 3 {
			return nil, errors.New("Input not kosher | " + s)
		}

		params = []string{params[0], params[1], params[2] + " " + params[3]}
	}

	//inputTime := fmt.Sprintf("%s %s", params[2], params[3])
	meetupTime, err := time.Parse("2006-01-02 3:04pm", params[2])
	if err != nil {
		return nil, err
	}

	meetupToInsert := &proto.Meetup{
		Name:     strings.TrimSpace(params[0]),
		Location: strings.TrimSpace(params[1]),
		Time:     meetupTime,
	}

	return meetupToInsert, nil
}
