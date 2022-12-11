run_on_local:
	go run main.go

google_pkey:
	http https://accounts.google.com/.well-known/openid-configuration

google_jwt_pkey:
	http https://www.googleapis.com/oauth2/v3/certs

build:
	docker compose build

run:
	docker compose up

rm:
	docker compose rm app -f
