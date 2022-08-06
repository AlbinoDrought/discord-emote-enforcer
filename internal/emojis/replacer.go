package emojis

import "strings"

var Replacer *strings.Replacer

func init() {
	emojiReplacerMux := make([]string, len(Emojis)*2)
	for i, emoji := range Emojis {
		emojiReplacerMux[i*2] = emoji
		emojiReplacerMux[(i*2)+1] = ""
	}
	Replacer = strings.NewReplacer(
		append(
			emojiReplacerMux,
			"\xEF\xB8\x8F", // variation selector
			"",
			"\xE2\x80\x8D", // ZWJ
			"",
		)...,
	)
	emojiReplacerMux = nil
}
