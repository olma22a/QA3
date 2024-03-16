package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var numericKeyboard1 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Да", "1"),
		tgbotapi.NewInlineKeyboardButtonData("Нет", "((("),
		tgbotapi.NewInlineKeyboardButtonData("Наверное", "пойдет"),
	),
)
var numericKeyboard2 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Да1", "5"),
		tgbotapi.NewInlineKeyboardButtonData("Нет1", "2"),
		tgbotapi.NewInlineKeyboardButtonData("Наверное1", "3"),
	),
)
var numericKeyboard3 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Да2", "1"),
		tgbotapi.NewInlineKeyboardButtonData("Нет2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("Наверное2", "3"),
	),
)

type Zalupa struct {
	Answer string `bson:"ответ"`
}

func main() {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("test").Collection("trainers")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы блинница?")
				msg.ReplyMarkup = numericKeyboard1
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
			}
		}

		if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {

			case "1":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Нажмите да1")
				msg.ReplyMarkup = numericKeyboard2
				var jopa Zalupa
				jopa.Answer = update.CallbackQuery.Data
				insertResult, err := collection.InsertOne(context.TODO(), jopa)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Inserted a single document: ", insertResult.InsertedID)
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}

			case "5":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Нажмите наверное2")
				msg.ReplyMarkup = numericKeyboard3
				var jopa Zalupa
				jopa.Answer = update.CallbackQuery.Data
				insertResult, err := collection.InsertOne(context.TODO(), jopa)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Inserted a single document: ", insertResult.InsertedID)
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}

			case "3":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "пшел пизду")
				var jopa Zalupa
				jopa.Answer = update.CallbackQuery.Data
				insertResult, err := collection.InsertOne(context.TODO(), jopa)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Inserted a single document: ", insertResult.InsertedID)
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
			}
		}
	}
}
