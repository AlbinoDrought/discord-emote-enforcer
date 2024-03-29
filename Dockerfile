FROM golang:1.18 as builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go get && CGO_ENABLED=0 go test && CGO_ENABLED=0 go build -o /discord-emote-channel

FROM alpine:3.14

RUN apk add --update --no-cache tini ca-certificates
USER 1000
COPY --from=builder /discord-emote-channel /discord-emote-channel
ENTRYPOINT ["tini", "--"]
CMD ["/discord-emote-channel"]
