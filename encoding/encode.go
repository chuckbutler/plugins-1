package encoding

import (
	"encoding/base64"
	"fmt"

	"github.com/go-chat-bot/bot"
)

const (
	invalidAmountOfParams = "Invalid amount of parameters"
	invalidParam          = "Invalid parameters"
)

func encode(command *bot.Cmd) (string, error) {

	if len(command.Args) != 2 {
		return invalidAmountOfParams, nil
	}

	var str string
	var err error
	switch command.Args[0] {
	case "base64":
		str, err = encodeBase64(command.Args[1])
	default:
		return invalidParam, nil
	}

	if err != nil {
		return fmt.Sprintf("Error: %s", err), nil
	}

	return str, nil
}

func encodeBase64(str string) (string, error) {
	data := []byte(str)
	return base64.StdEncoding.EncodeToString(data), nil
}

func init() {
	bot.RegisterCommand(
		"encode",
		"Allows you encoding a value",
		"base64 \"enter text to encode\"",
		encode)
}
