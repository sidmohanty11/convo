package controllers

import (
	"chat/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) CreateMessage(c echo.Context) (err error) {
	u := &models.User{
		ID: bson.ObjectIdHex(userIDFromToken(c)),
	}
	m := &models.Message{
		ID:   bson.NewObjectId(),
		From: u.ID.Hex(),
	}
	if err = c.Bind(m); err != nil {
		return
	}

	// Validation
	if m.To == "" || m.Message == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid to or message fields"}
	}

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
	userID := userIDFromToken(c)

	// Retrieve messages from database
	messages := []*models.Message{}
	db := h.DB.Clone()
	if err = db.DB("chat").C("messages").
		Find(bson.M{"from": userID}).
		All(&messages); err != nil {
		return
	}
	defer db.Close()

	return c.JSON(http.StatusOK, messages)
}
