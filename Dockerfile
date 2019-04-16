FROM mercari/appengine-go:1.11-alpine as builder

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
