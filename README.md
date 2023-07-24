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

connection string for goose:
host=localhost port=5432 user=rick password=picklerick dbname=GalacticFederation sslmode=disable
