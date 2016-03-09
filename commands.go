// Copyright 2016 The R2D2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/thoj/go-ircevent"
)

// event.Message() contains the message
// event.Nick Contains the sender
// event.Arguments[0] Contains the channel

func CommandHelp(ircobj *irc.Connection, event *irc.Event) {
	channel := event.Arguments[0]
	ircobj.Privmsg(channel, fmt.Sprintf("%s: The following commands are available:", event.Nick))
}
