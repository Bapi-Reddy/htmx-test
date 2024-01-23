FROM golang:1.21-alpine as build
RUN apk update && apk add --no-cache make git tzdata ca-certificates build-base && update-ca-certificates
WORKDIR /go/src/app
COPY . .
RUN touch todo.db 
RUN go mod tidy
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags='-w -s -extldflags "-static"' -v -o app main.go

# FROM scratch
# COPY --from=build /go/src/app/app /go/bin/
# COPY --from=build /go/src/app/todo.db /go/bin/
# COPY --from=build /go/src/app/css/ /go/bin/css/
# EXPOSE 3000
ENTRYPOINT ["/go/src/app/app"]