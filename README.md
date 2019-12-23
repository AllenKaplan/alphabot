# AlphaBot is a Discord bot in Golang

## How does it work?
* main.go starts the bot -> calls addHandler on the api handler
* only handler is the api handler -> handler calls services

## Current services
* Weather
    * `,weather get [location]` returns the weather from a given location based on the Open Weather Maps API
* Meetup
    * `,meetup create [name] [location] [YYYY-MM-DD] [HH:MM]` creates a meetup of the given params
    * `,meetup get [all|name]` gets all meetups or gets a meetup by specific name

## TODO
* Meetup
    * Add who created the meetup
    * `,meetup signup` registers a user for a meetup
    * `,meetup update [name] [location] [YYYY-MM-DD] [HH:MM]` allow meetup creator to update details
* Infrastructure
    * Add repo layer for data persistence
    * Deploy to Google App Engine    

### How to run?
`go run main.go -t [token]`

### How to build?
`go build .`