FROM golang:1.21.5 as build
RUN go install github.com/a-h/templ/cmd/templ@latest
WORKDIR /go/src/app
COPY . .
RUN touch /go/src/app/todo.db
ENV CGO_ENABLED=1 GOOS=linux GOPROXY=direct ARCH="amd64"
RUN templ generate && go build -ldflags='-w -s -extldflags "-static"' -v -o app .

FROM scratch
COPY --from=build /go/src/app/ /go/bin/
ENTRYPOINT ["/go/bin/app"]