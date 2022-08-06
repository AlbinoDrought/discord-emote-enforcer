package main

import (
	"fmt"
	"testing"
)

func Test_textContainsNonEmotes(t *testing.T) {
	tests := []struct {
		messageText string
		want        bool
	}{
		// animated emotes weren't being matched
		{
			"<a:snip:1234>",
			false,
		},
		{
			"<:snip:1234> <:snip:1234> <a:snip:1234> <:snip:1234> 💯<:snip:1234>",
			false,
		},
		// text basecase
		{
			"How did you",
			true,
		},
		// text + emoji basecase
		{
			"How did you 🏟️",
			true,
		},

		// real life below this point
		{
			"🕓 🚶 🏈 🏟️", // failing with go-emoji-utils
			false,
		},
		{
			"🕓 🏃 🏈 🏟️", // failing with go-emoji-utils
			false,
		},
		{
			"🕓 🚗 🏈 🏟️",
			false,
		},
		{
			"👨‍💼👨‍💻🥲", // failing with go-emoji-utils + stock data, passes with aftermarket
			false,
		},
		{
			"🧑‍💼🥲", // failing with go-emoji-utils + stock data, passes with aftermarket
			false,
		},
		{
			"🧑‍💻🥲", // failing with go-emoji-utils + stock data, passes with aftermarket
			false,
		},
		{
			"🧑‍💼🥲", // failing with go-emoji-utils + stock data, passes with aftermarket
			false,
		},
		{
			"☁️, 🆗️❓", // contains comma
			true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%v", i), func(t *testing.T) {
			if got := textContainsNonEmotes(tt.messageText); got != tt.want {
				t.Errorf("textContainsNonEmotes(%v) = %v, want %v, leftovers %v, %x", tt.messageText, got, tt.want, removeEmotes(tt.messageText), removeEmotes(tt.messageText))
			}
		})
	}
}
