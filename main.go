package main

import (
	"regexp"
)

// ConfigFile is the path of the default configuration file
var ConfigFile = "config.json"

// Gateway
var IPFSGateway = "https://ipfs.io"

const twittercom = "twitter.com"

type Action string

// Variables containing the different available actions
var (
	// (spaces)(action)whitespaces(arguments)
	actionRegexp = regexp.MustCompile(`^\s*([[:graph:]]+)\s+(.+)`)
	// (cid)whitespaces(name with whitespaces). [:graph:] does not
	// match line breaks or spaces.
	pinRegexp          = regexp.MustCompile(`([[:graph:]]+)\s+([[:graph:]\s]+)`)
	PinAction   Action = "!pin"
	UnpinAction Action = "!unpin"
	AddAction   Action = "!add"
	HelpAction  Action = "!help"
)

func main() {
	//Let's code ☕
}
