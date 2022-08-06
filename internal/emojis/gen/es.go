package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	btxt, err := ioutil.ReadFile("emoji-style.txt") // https://unicode.org/emoji/charts/emoji-style.txt
	if err != nil {
		panic(err)
	}
	txt := string(btxt)
	pr := regexp.MustCompile("\n(?:text|emoji|modifier|zwj).*\n((.+\n)+)")

	rawTxt := ""

	for _, match := range pr.FindAllStringSubmatch(txt, -1) {
		rawTxt += match[1]
	}

	txt = strings.ReplaceAll(rawTxt, "\n", " ")
	txt = strings.TrimSpace(txt)

	emojis := strings.Split(txt, " ")
	j, err := json.Marshal(emojis)
	if err != nil {
		panic(err)
	}

	j[0] = '{'
	j[len(j)-1] = '}'

	data := "package emojis\n\nvar Emojis = []string" + string(j)
	ioutil.WriteFile("emojis.go", []byte(data), os.ModePerm)
}
