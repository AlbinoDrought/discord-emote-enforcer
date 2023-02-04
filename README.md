# Emote Enforcer

Deletes new messages in a Discord channel if they contain something other than emojis or Discord emotes.

## Usage

```sh
DEC_TOKEN=your-discord-bot-token \
DEC_GUILD_ID=your-guild-id \
DEC_CHANNEL_ID=your-channel-id,your-other-channel-id \
go run main.go
```

See https://discord.com/developers/docs/topics/oauth2#bots for information on creating a Discord bot.
