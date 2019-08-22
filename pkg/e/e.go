package e

import "github.com/tidwall/gjson"

type MarketError struct {
	Message string `json:"message"`
}

func GetMsg(body string) string {
	msg := gjson.Get(body, "message")
	return msg.String()
}
