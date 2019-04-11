FROM mercari/appengine-go:1.11-alpine as builder

COPY . .

RUN go get
