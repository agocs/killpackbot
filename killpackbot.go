package main

import (
	"github.com/thoj/go-ircevent"
	"strings"
)

func main() {
	ircchan := "#pugt"
	ircobj := irc.IRC("killpackbot", "killpackbot")
	err := ircobj.Connect("irc.freenode.net:6667")
	if err != nil {
		print(err.Error())
	}
	ircobj.Join(ircchan)
	ircobj.AddCallback("PRIVMSG", func(e *irc.Event) {
		if strings.Contains(e.Message(), "wank") {
			wanker := e.Nick
			ircobj.Privmsgf(ircchan, "*frowns at %s*", wanker)
		}
		if strings.Contains(e.Message(), "killpackbot: help") {
			ircobj.Privmsgf(ircchan, "%s: Frowns at wankers", e.Nick)
		}
	})

	ircobj.Loop()
}
