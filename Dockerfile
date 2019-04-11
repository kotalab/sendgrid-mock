FROM mercari/appengine-go:1.11-alpine as builder

ENV PROJECT_ROOT /go/src/github.com/kotalab/sendgrid-mock/
WORKDIR $PROJECT_ROOT

RUN go get -u github.com/golang/dep/cmd/dep

# restore dependencies
COPY Gopkg.* ./
RUN dep ensure --vendor-only -v

COPY . .
