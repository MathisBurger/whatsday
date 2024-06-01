package main

import (
	"context"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
	"time"
	"whatsday/internal"
)

func main() {
	client := internal.ConnectToAPI()
	var config internal.FullConfig
	config.GetConfig()
	idMap := config.GetJIDMap(client)

	for {
		now := time.Now()
		formatted := now.Format("01-02")
		val, ok := idMap[formatted]
		if ok {
			for _, jid := range val {
				_, _ = client.SendMessage(context.Background(), jid, &waProto.Message{
					Conversation: proto.String(config.GetRandomMessage()),
				})
			}
		}
		time.Sleep(24 * time.Hour)
	}
}
