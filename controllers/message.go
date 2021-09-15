package controllers

import (
	"chat/models"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/pusher/pusher-http-go"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) GetChat(c echo.Context) (err error) {
	user := &models.User{}
	usernameTo := c.Param("username")
	usernameFrom := c.Param("userFrom")

	messages := []*models.Message{}
	messagesTo := []*models.Message{}
	db := h.DB.Clone()
	if err = db.DB("chat").
		C("users").
		Find(bson.M{"username": usernameTo}).One(&user); err != nil {
		return
	}

	if err = db.DB("chat").C("messages").
		Find(bson.M{"to": usernameFrom, "from": usernameTo}).
		All(&messagesTo); err != nil {
		return
	}

	if err = db.DB("chat").C("messages").
		Find(bson.M{"to": usernameTo, "from": usernameFrom}).
		All(&messages); err != nil {
		return
	}
	defer db.Close()

	return c.JSON(http.StatusOK, echo.Map{"messages_from": messages, "messagesTo": messagesTo, "user": user})
}

func (h *Handler) CreateMessage(c echo.Context) (err error) {
	u := &models.User{
		ID: bson.ObjectIdHex(userIDFromToken(c)),
	}
	m := &models.Message{
		ID:        bson.NewObjectId(),
		Timestamp: bson.MongoTimestamp(bson.Now().Day()),
	}
	if err = c.Bind(m); err != nil {
		return
	}

	// Validation
	if m.From == "" || m.To == "" || m.Message == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid to or message fields"}
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pusherClient := pusher.Client{
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_APP_KEY"),
		Secret:  os.Getenv("PUSHER_APP_SECRET"),
		Cluster: "ap2",
		Secure:  true,
	}

	data := echo.Map{"message": m}
	pusherClient.Trigger("my-channel", "my-event", data)

	// Find user from database
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("chat").C("users").FindId(u.ID).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
		return
	}

	// Save message in database
	if err = db.DB("chat").C("messages").Insert(m); err != nil {
		return
	}
	return c.JSON(http.StatusCreated, m)
}

func (h *Handler) FetchMessage(c echo.Context) (err error) {
	username := c.Param("username")
	db := h.DB.Clone()
	// Retrieve messages from database
	messages := []*models.Message{}
	if err = db.DB("chat").C("messages").
		Find(bson.M{"from": username}).
		All(&messages); err != nil {
		return
	}

	defer db.Close()

	return c.JSON(http.StatusOK, messages)
}
