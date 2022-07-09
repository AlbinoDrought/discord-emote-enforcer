FROM golang:1.18 as builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go get && CGO_ENABLED=0 go build -o /discord-emote-channel

FROM scratch
USER 1000
COPY --from=builder /discord-emote-channel /discord-emote-channel
CMD ["/discord-emote-channel"]
