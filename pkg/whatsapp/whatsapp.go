package whatsapp

import (
	"0day-backend/pkg/logging"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendWA(chatID string, Name string) {
	url := "http://whatsapp.findanime.to/api/sendText"
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		"X-Api-Key":    "",
	}

	data := map[string]interface{}{
		"chatId": fmt.Sprintf("91%s@s.whatsapp.net", chatID),
		"text": fmt.Sprintf(`Hey %s 🎉 *Welcome to Club 0day!* 🎉

You're officially a member! 🚀
We're thrilled to have you with us. Here's how you can stay connected:

💬 *Join our community:*
- WhatsApp: https://chat.whatsapp.com/DxVzQ3XbMdk6KMsmapmHN9

- Discord: https://discord.gg/VB4FMbuukE

📞 *Need help?* Reach out to Arya: 7603061337

*Note : If link is not clickable,forward this message to this account itself and it will work.*

Let's innovate, learn, and grow together! 💡
Excited to see you in action! 🔥`, Name), // <- pass the name variable here
		"session": "default",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		logging.Logger.Error().Str("Phone", chatID).Str("Name", Name).Msg(err.Error())
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		logging.Logger.Error().Str("Phone", chatID).Str("Name", Name).Msg(err.Error())
		return
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logging.Logger.Error().Str("Phone", chatID).Str("Name", Name).Msg(err.Error())

		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 201:
		logging.Logger.Info().Str("Phone", chatID).Str("Name", Name).Msg("Whatsapp DM Sent Successfully")
	default:
		logging.Logger.Error().Str("Phone", chatID).Str("Name", Name).Msg("Failed Whatsapp DM Sent Successfully")
	}
}
