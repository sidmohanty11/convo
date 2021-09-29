package handlers

import (
	"convo_backend/models"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
)

var (
	upgrader = websocket.Upgrader{}
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
