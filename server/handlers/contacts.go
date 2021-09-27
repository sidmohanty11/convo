package handlers

import (
	"convo_backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) AddContactOfUser(c echo.Context) (err error) {
	u := &models.User{
		ID: bson.ObjectIdHex(userIDFromToken(c)),
	}

	contact := &models.Contact{
		ID: bson.NewObjectId(),
	}
	if err = c.Bind(contact); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "something went wrong"}
	}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("convo").C("users").FindId(u.ID).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "db not responding"}
	}

	u.Contacts = append(u.Contacts, contact)

	db.DB("convo").C("users").UpdateId(u.ID, u)

	return c.JSON(http.StatusOK, contact)
}
