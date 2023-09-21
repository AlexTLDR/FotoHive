# Production Stuff

I am still waiting on mailtrap to allow me to send emails. Waiting on them to validate the email sending domain.I need this validation for password recovery emails.

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
