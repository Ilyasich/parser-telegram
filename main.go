package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"telegram-vacancy-parser/parser"
	"telegram-vacancy-parser/storage"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from environment variables")
	}

	appIDStr := os.Getenv("APP_ID")
	appHash := os.Getenv("APP_HASH")
	phone := os.Getenv("PHONE")

	if appIDStr == "" || appHash == "" {
		log.Fatal("APP_ID and APP_HASH must be set")
	}

	appID, err := strconv.Atoi(appIDStr)
	if err != nil {
		log.Fatalf("Invalid APP_ID: %v", err)
	}

	// Helper for terminal authentication
	flow := auth.NewFlow(
		auth.Constant(phone, "", auth.CodeAuthenticatorFunc(func(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
			fmt.Print("Enter code: ")
			var code string
			_, err := fmt.Scan(&code)
			return code, err
		})),
		auth.SendCodeOptions{},
	)

	// Initialize storage
	store := storage.NewFileStorage("vacancies.json")

	// Dispatcher handles incoming updates
	dispatcher := tg.NewUpdateDispatcher()
	dispatcher.OnNewMessage(func(ctx context.Context, e tg.Entities, update *tg.UpdateNewMessage) error {
		handleMessage(update.Message, store)
		return nil
	})
	dispatcher.OnNewChannelMessage(func(ctx context.Context, e tg.Entities, update *tg.UpdateNewChannelMessage) error {
		handleMessage(update.Message, store)
		return nil
	})

	client := telegram.NewClient(appID, appHash, telegram.Options{
		UpdateHandler: dispatcher,
	})

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := client.Run(ctx, func(ctx context.Context) error {
		if err := client.Auth().IfNecessary(ctx, flow); err != nil {
			return err
		}

		user, err := client.Self(ctx)
		if err != nil {
			return err
		}
		log.Printf("Logged in as %s (%s)\n", user.FirstName, user.Username)
		log.Println("Listening for vacancies...")

		<-ctx.Done()
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func handleMessage(msgClass tg.MessageClass, store *storage.FileStorage) {
	switch m := msgClass.(type) {
	case *tg.Message:
		if parser.IsVacancy(m.Message) {

			peerID := getPeerID(m.PeerID)
			log.Printf("[VACANCY FOUND] ChatID: %d | Content: %s\n", peerID, m.Message)

			v := storage.Vacancy{
				ChatID:    peerID,
				MessageID: m.ID,
				Text:      m.Message,
				FoundAt:   time.Now(),
			}

			if err := store.Save(v); err != nil {
				log.Printf("Failed to save vacancy: %v\n", err)
			}
		}
	}
}

func getPeerID(peer tg.PeerClass) int64 {
	switch p := peer.(type) {
	case *tg.PeerUser:
		return p.UserID
	case *tg.PeerChat:
		return p.ChatID
	case *tg.PeerChannel:
		return p.ChannelID
	}
	return 0
}
