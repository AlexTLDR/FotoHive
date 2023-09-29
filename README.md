# Production Stuff

I am currently awaiting Mailtrap to grant me permission to send emails. I'm waiting for them to verify the email sending domain. This validation is necessary for sending password recovery emails.

https://fotohive.alextldr.com/


# WebDev

A place to hack/test WebDev stuff with Go

# Docker & Postgres

$docker compose up

Env:
Server: db
POSTGRES_USER: rick
POSTGRES_PASSWORD: picklerick
POSTGRES_DB: GalacticFederation
Adminer:
localhost:3333
PSQL:
docker compose exec -it db psql -U rick -d GalacticFederation

# connection string for goose:

host=localhost port=5432 user=rick password=picklerick dbname=GalacticFederation sslmode=disable

# Goose install:

go install github.com/pressly/goose/v3/cmd/goose@latest
export GOPATH=$HOME/alex/git
export PATH=$PATH:/home/alex/go:$GOPATH/bin
