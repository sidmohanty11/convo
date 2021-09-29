package handlers

import (
	"convo_backend/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/pusher/pusher-http-go"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) MessageByChatId(c echo.Context) (err error) {
	id := c.Param("id")

	chat := &models.Chat{}

	msg := &models.Message{
		ID:        bson.NewObjectId(),
		Timestamp: time.Now(),
	}

	if err = c.Bind(msg); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}
	}

	err = godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	pusherClient := pusher.Client{
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_APP_KEY"),
		Secret:  os.Getenv("PUSHER_APP_SECRET"),
		Cluster: "ap2",
		Secure:  true,
	}

	data := echo.Map{"message": msg}
	pusherClient.Trigger("my-channel", "my-event", data)

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("convo").C("chats").FindId(bson.ObjectIdHex(id)).One(chat); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}
	}

	chat.Messages = append(chat.Messages, msg)

	if err = db.DB("convo").C("chats").UpdateId(chat.ID, chat); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}
	}

	return c.JSON(http.StatusOK, msg)
}
