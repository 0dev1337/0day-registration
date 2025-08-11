package main

import (
	"0day-backend/internal/routes"
	"fmt"
	"log"
	"time"

	whatsapp "github.com/Rhymen/go-whatsapp"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	wac, _ := whatsapp.NewConn(20 * time.Second)
	qrChan := make(chan string)
	go func() {
		fmt.Printf("qr code: %v\n", <-qrChan)
	}()
	sess, err := wac.Login(qrChan)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(sess.ClientToken)

	app := fiber.New()
	app.Use(cors.New())

	// Setup Routes
	routes.SetupRoutes(app)
	app.Listen(":3001")
}
