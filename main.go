package main

import (
	"os"

	"errors"
	"fmt"
	"log"
	"time"

	"github.com/alash3al/go-smtpsrv"
	"github.com/go-resty/resty/v2"
)

func main() {
	cfg := smtpsrv.ServerConfig{
		ReadTimeout:     time.Duration(*flagReadTimeout) * time.Second,
		WriteTimeout:    time.Duration(*flagWriteTimeout) * time.Second,
		ListenAddr:      *flagListenAddr,
		MaxMessageBytes: int(*flagMaxMessageSize),
		BannerDomain:    *flagServerName,
		Handler: smtpsrv.HandlerFunc(func(c *smtpsrv.Context) error {
			msg, err := c.Parse()
			if err != nil {
				return errors.New("Unable to parse the message: " + err.Error())
			}

			// Read environment variables
			apiURL := os.Getenv("ZOHO_API_URL")
			authToken := os.Getenv("ZOHO_TOKEN")
			fromAddress := os.Getenv("ZOHO_FROM_ADDRESS")

			if apiURL == "" || authToken == "" || fromAddress == "" {
				return errors.New("Environment variables ZOHO_API_URL, ZOHO_TOKEN, and ZOHO_FROM_ADDRESS are required")
			}

			to := c.To().Address
			subject := msg.Subject

			// Prefer plain text body, fallback to HTML
			body := string(msg.TextBody)
			if body == "" {
				body = string(msg.HTMLBody)
			}

			// Build the payload for Zoho Mail API
			payload := map[string]string{
				"fromAddress": fromAddress,
				"toAddress":   to,
				"subject":     subject,
				"content":     body,
			}

			// Send POST request
			resp, err := resty.New().
				R().
				SetHeader("Authorization", authToken).
				SetHeader("Content-Type", "application/json").
				SetBody(payload).
				Post(apiURL)

			if err != nil {
				log.Println("Error while sending request to Zoho Mail API:", err)
				return errors.New("Internal error while processing the message (E1)")
			}

			if resp.StatusCode() >= 300 {
				log.Printf("Zoho API response: %d - %s", resp.StatusCode(), resp.String())
				return fmt.Errorf("Zoho Mail API returned an error: %s", resp.String())
			}

			log.Printf("Email successfully forwarded to Zoho: %s", to)
			return nil
		}),
	}

	fmt.Println(smtpsrv.ListenAndServe(&cfg))
}
