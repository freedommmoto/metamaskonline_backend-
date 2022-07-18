package tool

import (
	"github.com/pusher/pusher-http-go/v5"
	"log"
)

func CallPusherClient(idUser string, PusherKey string) {

	log.Println("in function CallPusherClient")

	pusherClient := pusher.Client{
		AppID:   "1431028",
		Key:     "3b3bee31bf863d7fa58d",
		Secret:  PusherKey,
		Cluster: "ap1",
	}
	data := map[string]string{"idUser is ": idUser}

	// trigger an event on a channel, along with a data payload
	err := pusherClient.Trigger("channel-userid-"+idUser, "code-active", data)

	// All trigger methods return an error object, it's worth at least logging this!
	if err != nil {
		panic(err)
	}
}
