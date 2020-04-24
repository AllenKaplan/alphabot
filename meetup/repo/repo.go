package repository

import (
	"fmt"
	"github.com/AllenKaplan/alphabot/meetup/proto"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // here
	"os"
)

type MeetupRepo struct{
	databaseType string
	connectionString string
}

func NewMeetupRepo(databaseType string, connectionString string) *MeetupRepo {
	return &MeetupRepo{
		databaseType:     databaseType,
		connectionString: connectionString,
	}
}


func CreateConnection(tableName string, connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open(tableName, connectionString)

	return db, err
}

func (repo *MeetupRepo) CreateMeetup(*meetup.Meetup) (bool, error) {
	panic("implement me")
}

func (repo *MeetupRepo) GetMeetup(string) (*meetup.Meetup, error) {
	db, err := CreateConnection(repo.databaseType, repo.connectionString)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	var queryResult []*quoteModel

	if err := db.Table("quotes").Where("user_id = ?", id).Find(&queryResult).Error; err != nil {
		//returns false if there is an error but the Quote might actually exist
		return nil, err
	}

	result := quoteModelsToQuoteEntries(queryResult)

	return result, nil
}

func (repo *MeetupRepo) GetAllMeetups() ([]*meetup.Meetup, error) {
	panic("implement me")
}

func InitRepo() *MeetupRepo {
	// Get database details from environment variables
	host := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	if host == "" { host = "localhost"}
	if dbUser == "" { dbUser = "postgres" }
	if password == "" { password = "password"}
	if DBName == "" { DBName = "quoter" }

	connectionString := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		host, dbUser, DBName, password,
	)

	databaseType := "postgres"

	return NewMeetupRepo(databaseType, connectionString)
}