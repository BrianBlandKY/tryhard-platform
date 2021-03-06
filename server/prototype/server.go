package prototype

import (
	"github.com/labstack/echo"
)

// Server Interface
type Server interface {
	Add(Client)
	Drop(Client)
	Error(error)
	Write(Message)
	Done()
	Listen(*echo.Echo)
}
