// Copyright 2016 The R2D2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	inidata, err := ioutil.ReadFile("r2d2.ini")
	if err != nil {
		log.Fatalln("Failed to load r2d2.ini:", err)
		return
	}

	cfg, err := ini.Load(inidata)
	if err != nil {
		log.Fatalln("Failed to parse configuration:", err)
		return
	}

	section, err := cfg.GetSection("Connection")
	if err != nil {
		log.Fatalln("Failed to get connection information:", err)
		return
	}

	channel := fmt.Sprintf("#%s", section.Key("Channel").String())

	ircobj := irc.IRC(section.Key("Nick").String(), section.Key("User").String())
	ircobj.UseTLS = section.Key("UseTLS").MustBool(false)
	//ircobj.VerboseCallbackHandler = true
	//ircobj.Debug = true
	if section.HasKey("ServerPassword") {
		ircobj.Password = section.Key("ServerPassword").String()
	}

	ircobj.AddCallback("001", func(event *irc.Event) {
		ircobj.Join(channel)
	})

	ircobj.AddCallback("366", func(event *irc.Event) {
		//ircobj.Privmsg(channel, "beep beep beep\n")
	})

	ircobj.AddCallback("PRIVMSG", func(event *irc.Event) {
		go func(event *irc.Event) {
			message := event.Message()

			if strings.HasPrefix(message, "!help") {
				CommandHelp(ircobj, event)
			} else if strings.HasPrefix(message, "!google") {
				args := strings.Split(strings.Replace(message, "!google ", "", 1), " ")
				CommandGoogle(ircobj, event, args)
			}
		}(event)
	})

	err = ircobj.Connect(section.Key("Server").String())
	if err != nil {
		log.Fatalln(err)
		return
	}

	ircobj.Loop()
}
