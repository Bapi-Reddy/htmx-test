FROM golang:1.21.5-bookworm AS build
WORKDIR /go/src/app
COPY . .
RUN touch /go/src/app/todo.db
ENV CGO_ENABLED=1 GOOS=linux GOPROXY=direct ARCH="amd64"
RUN go build -ldflags='-w -s -extldflags "-static"' -v -o app .

FROM scratch
COPY --from=build /go/src/app/app /go/bin/app
COPY --from=build /go/src/app/todo.db /go/bin/todo.db
ENTRYPOINT ["/go/bin/app"]