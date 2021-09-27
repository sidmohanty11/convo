package handlers

import (
	"convo_backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) AddChat(c echo.Context) (err error) {
	chat := &models.Chat{
		ID: bson.NewObjectId(),
	}

	if err = c.Bind(chat); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "something went wrong"}
	}

	db := h.DB.Clone()
	defer db.Close()

	for _, person := range chat.Participants {
		u := &models.User{}
		if err = db.DB("convo").C("users").FindId(bson.ObjectIdHex(person)).One(u); err != nil {
			if err == mgo.ErrNotFound {
				return echo.ErrNotFound
			}
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "db not responding"}
		}

		u.Chats = append(u.Chats, chat)

		if err = db.DB("convo").C("users").UpdateId(bson.ObjectIdHex(person), u); err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "db not responding"}
		}
	}

	if err = db.DB("convo").C("chats").Insert(chat); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "db not responding"}
	}

	return c.JSON(http.StatusOK, chat)
}

func (h *Handler) GetChatById(c echo.Context) (err error) {
	id := c.Param("id")

	chat := &models.Chat{}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("convo").C("chats").FindId(bson.ObjectIdHex(id)).One(chat); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	return c.JSON(http.StatusOK, chat)
}
