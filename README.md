# AlphaBot is your source for a well written discord go bot

## How does it work?
* main.go starts the bot -> calls addHandler on the api handler
* only handler is the api handler -> handler calls services

## Current services
* weather

### How to run?
`go run main.go -t [token]`

### How to build?
`go build .`