# Production Stuff

https://fotohive.alextldr.com/

This is a photo sharing app that I created in order to showcase my Go skills. It uses PostgreSQL for the database, Go for the backend, and gohtml templates for the front end.

When registering a new account, I bypassed the email validation, allowing the use of any mock email address. However, to test the "Forgot password?" feature, a valid email address is necessary.

You can utilize https://10minutemail.net/ to generate a temporary email address for this purpose.


Photos can be uploaded from your local machine, or via Dropbox.


I've set a maximum file size of 10MB, and I'm also checking that files are in either the gif, jpeg, or png formats. Attempting to upload a file in a different format won't be feasible.


I'm enhancing security by encrypting passwords in the database using both hashes and salts. This means that passwords are not visible in the database.

Additionally, because of the unique salts, even if two different accounts use the same password, the hashes stored in the database will be distinct. To prevent session hijacking, I'm employing session tokens.

# TODO

Enhance error handling - substitute generic error messages with specific ones.

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
