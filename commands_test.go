// Copyright 2016 The R2D2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCommandGoogle(t *testing.T) {
	args := []string{"search", "this"}
	sender := "nick"
	output := CommandGoogle("#chan", sender, args)

	query := strings.Join(args, "+")
	uri := fmt.Sprintf("http://lmgtfy.com/?q=%s", url.QueryEscape(query))
	if output[0] != fmt.Sprintf("%s: %s", sender, uri) {
		t.Fatal("url doesn't match")
	}
}
