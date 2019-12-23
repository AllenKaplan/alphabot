package meetup

import (
	"fmt"
	"time"
)

type Meetup struct {
	Name     string
	Location string
	Time     time.Time
}

func (m Meetup) String() string {
	return fmt.Sprintf("%s | %s | %s", m.Name, m.Location, m.Time.Format("2006-01-02"))
}
