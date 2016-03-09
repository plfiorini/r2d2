// Copyright 2016 The R2D2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/url"
	"strings"
)

func CommandHelp(channel, sender string) []string {
	output := make([]string, 0)
	output = append(output, fmt.Sprintf("%s: The following commands are available:", sender))
	return output
}

func CommandGoogle(channel string, sender string, args []string) []string {
	output := make([]string, 0)
	query := strings.Join(args, "+")
	uri := fmt.Sprintf("http://lmgtfy.com/?q=%s", url.QueryEscape(query))
	output = append(output, fmt.Sprintf("%s: %s", sender, uri))
	return output
}
