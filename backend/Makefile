build:
		GOOS=linux go build -o app
		docker build -t gcr.io/golang-stuffz/say .
		rm -f app

push:
		gcloud auth configure-docker grc.io
		docker -- push gcr.io/golang-stuffz/say