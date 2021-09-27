package routes

import (
	"convo_backend/handlers"

	"github.com/labstack/echo/v4"
)

func Setup(h *handlers.Handler, e *echo.Echo) {
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)

	e.GET("/activeUser", h.GetActiveUser)
	e.GET("/user/:id", h.GetUserById)
	e.GET("/users/:number", h.GetUserByNumber)

	e.POST("/contacts/add", h.AddContactOfUser)

	e.POST("/chat/add", h.AddChat)
	e.GET("/chats/:id", h.GetChatById)

	e.POST("/messages/add/:id", h.MessageByChatId)
}
