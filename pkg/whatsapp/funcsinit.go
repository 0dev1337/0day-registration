package whatsapp

// import (
// 	"encoding/gob"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"time"

// 	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
// 	"github.com/Rhymen/go-whatsapp"
// )

// type WhatsAppClient struct {
// 	conn *whatsapp.Conn
// }

// // NewClient creates a new WhatsApp client and logs in or restores session.
// func NewClient() (*WhatsAppClient, error) {
// 	conn, err := whatsapp.NewConn(20 * time.Second)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create WhatsApp connection: %w", err)
// 	}

// 	client := &WhatsAppClient{conn: conn}

// 	if err := client.login(); err != nil {
// 		return nil, err
// 	}

// 	return client, nil
// }

// // SendText sends a text message to the given phone number in international format (no +).
// func (wac *WhatsAppClient) SendText(phoneNumber, message string) error {
// 	jid := fmt.Sprintf("%s@s.whatsapp.net", phoneNumber)

// 	msg := whatsapp.TextMessage{
// 		Info: whatsapp.MessageInfo{
// 			RemoteJid: jid,
// 		},
// 		Text: message,
// 	}

// 	msgID, err := wac.conn.Send(msg)
// 	if err != nil {
// 		return fmt.Errorf("failed to send message: %w", err)
// 	}

// 	fmt.Println("Message sent successfully. ID:", msgID)
// 	return nil
// }

// // login tries to restore a previous session or logs in via QR code.
// func (wac *WhatsAppClient) login() error {
// 	session, err := readSession()
// 	if err == nil {
// 		// Try to restore session
// 		session, err = wac.conn.RestoreWithSession(session)
// 		if err != nil {
// 			fmt.Println("Failed to restore session, logging in with QR code.")
// 		} else {
// 			fmt.Println("Session restored successfully.")
// 			return nil
// 		}
// 	}

// 	// Login via QR code
// 	qr := make(chan string)
// 	go func() {
// 		terminal := qrcodeTerminal.New()
// 		terminal.Get(<-qr).Print()
// 	}()

// 	session, err = wac.conn.Login(qr)
// 	if err != nil {
// 		return fmt.Errorf("failed to login via QR code: %w", err)
// 	}

// 	if err := writeSession(session); err != nil {
// 		fmt.Println("Warning: failed to save session:", err)
// 	}
// 	fmt.Println("Logged in successfully.")
// 	return nil
// }

// // readSession loads session from file.
// func readSession() (whatsapp.Session, error) {
// 	session := whatsapp.Session{}
// 	path := sessionFilePath()

// 	file, err := os.Open(path)
// 	if err != nil {
// 		return session, err
// 	}
// 	defer file.Close()

// 	decoder := gob.NewDecoder(file)
// 	err = decoder.Decode(&session)
// 	return session, err
// }

// // writeSession saves session to file.
// func writeSession(session whatsapp.Session) error {
// 	path := sessionFilePath()

// 	file, err := os.Create(path)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	encoder := gob.NewEncoder(file)
// 	return encoder.Encode(session)
// }

// // sessionFilePath returns the full path to the session file.
// func sessionFilePath() string {
// 	return filepath.Join(os.TempDir(), "whatsappSession.gob")
// }
