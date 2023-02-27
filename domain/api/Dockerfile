FROM golang:1.20-alpine as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN apk update

RUN mkdir -p /tmp/labyrinth

WORKDIR /tmp/labyrinth
ADD domain .

RUN go build -ldflags="-w -s" -o /tmp/app ./api

FROM scratch

COPY --from=builder /tmp/app /app

EXPOSE 80

ENTRYPOINT ["/app"]
