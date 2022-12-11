run:
	go run main.go

google_pkey:
	http https://accounts.google.com/.well-known/openid-configuration

google_jwt_pkey:
	http https://www.googleapis.com/oauth2/v3/certs
