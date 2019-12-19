# AlphaBot is a Discord bot in Golang

## How does it work?
* main.go starts the bot -> calls addHandler on the api handler
* only handler is the api handler -> handler calls services

## Current services
* weather

### How to run?
`go run main.go -t [token]`

### How to build?
`go build .`