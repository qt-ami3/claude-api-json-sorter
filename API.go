package main

import (
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

var client = anthropic.NewClient (
	option.WithAPIKey(""), // defaults to os.LookupEnv("ANTHROPIC_API_KEY")
)

var modelNameAPI = anthropic.ModelClaudeHaiku4_5_20251001
