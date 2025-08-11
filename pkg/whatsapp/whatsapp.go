package whatsapp

import (
	"0day-backend/pkg/logging"
	"fmt"
	"log"
)

var (
	Client *WhatsAppClient
)

// func init() {
// 	var err error
// 	Client, err = NewClient()
// 	if err != nil {
// 		logging.Logger.Error().Msg(err.Error())
// 	}

// }

func SendDM(number string) {

	err := Client.SendText(fmt.Sprintf("+91%s", number), "Hello from Go WhatsApp client!")
	if err != nil {
		log.Fatalln("Failed to send message:", err)
	}
	logging.Logger.Info().Str("phone", number).Msg("Successfully Send DM Message To")
}
