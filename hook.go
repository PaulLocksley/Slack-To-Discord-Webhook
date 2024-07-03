package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// WebhookPayload represents the structure of your webhook payload
type WebhookPayload struct {
	Text string `json:"text"`
}

type DiscordMessege struct {
	Content string `json:"content"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	var payload WebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = discordMessenger(DiscordMessege{payload.Text})
	if err != nil {
		http.Error(w, "Failed to forward message", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received successfully"))
}

func discordMessenger(message DiscordMessege) error {
	webhookUrl := os.Getenv("s2dwebhook")
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return err
	}

	resp, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(jsonMessage))
	if err != nil {
		fmt.Printf("Error sending HTTP request: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected response status: %v", resp.Status)
	}

	fmt.Println("Message sent successfully!")
	return nil
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)

	fmt.Println("Listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
