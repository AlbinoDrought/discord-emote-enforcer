package main

import (
	_ "embed"
	"encoding/json"

	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

//go:embed data/emoji.json
var emojiData []byte

func init() {
	if err := json.Unmarshal(emojiData, &emoji.Emojis); err != nil {
		panic(err)
	}
}

func RemoveAll(text string) string {
	return emoji.RemoveAll(text)
}
