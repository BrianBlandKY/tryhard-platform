package server

import (
	_ "fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/websocket"
	// std "github.com/labstack/echo/engine/standard"
)

var CLIENT_KEY string = "CLIENT_ID"
var CLIENT_EXPIRE_TIME_HOURS int = 24

type Server struct {
	origin      string
	partySource *partySource
	clients     map[string]*client
	addCh       chan *client
	delCh       chan *client
	doneCh      chan bool
	errCh       chan error
	messageCh   chan *message
}

func (s *Server) Listen(router *echo.Echo) {
	log.Println("Listening server...")

	// Make configurable.
	// Add option to disable/skip?
	origin, _ := url.ParseRequestURI(s.origin)
	ws := websocket.Server{
		Config: websocket.Config{
			Origin: origin,
		},
		Handler:   s.socketHandler(),
		Handshake: s.socketHandshake(),
	}
	router.Use(s.socketMiddleware)
	router.GET("/connect", echo.WrapHandler(ws))
	router.GET("/party/gen", s.partyGenCodeHandler())
	router.GET("/party/gen/:code", s.partyGenCodeHandler())
	router.GET("/party/:code", s.partyCodeHandler())
	log.Println("Socket Server Listening...")
	s.messageHandler()
}
func (s *Server) add(c *client) {
	s.addCh <- c
}
func (s *Server) drop(c *client) {
	s.delCh <- c
}
func (s *Server) done() {
	s.doneCh <- true
}
func (s *Server) err(err error) {
	s.errCh <- err
}
func (s *Server) write(msg *message) {
	s.messageCh <- msg
}
func (s *Server) socketMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientId := ""
		setNewCookie := true
		reqCookies := c.Cookies()

		for _, cookie := range reqCookies {
			if cookie.Name == CLIENT_KEY {
				clientId = cookie.Value
				setNewCookie = false
			}
		}
		if setNewCookie {
			clientId = uuid.NewV4().String()
			newCookie := new(http.Cookie)
			newCookie.Name = CLIENT_KEY
			newCookie.Value = clientId
			newCookie.Expires = time.Now().Add(time.Hour * time.Duration(CLIENT_EXPIRE_TIME_HOURS))
			c.SetCookie(newCookie)
		}

		// Add as Header to the request
		// This will be picked up in the handshake and passed to the socket Client.
		c.Request().Header.Add(CLIENT_KEY, clientId)

		return next(c)
	}
}
func (s *Server) socketHandler() websocket.Handler {
	return func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()

		clientID := ws.Config().Header.Get(CLIENT_KEY)
		if clientID == "" {
			log.Fatal("invalid clientid")
		}
		client := newClient(clientID, ws, s)
		s.add(client)
		client.listen()
	}
}
func (s *Server) socketHandshake() func(*websocket.Config, *http.Request) error {
	return func(config *websocket.Config, req *http.Request) (err error) {
		// Origin Validation
		// origin, err := websocket.Origin(config, req)
		// if err != nil {
		// 	return err
		// }
		// if err == nil && origin == nil {
		// 	return fmt.Errorf("null origin")
		// }

		// if origin.String() != config.Origin.String() {
		// 	return fmt.Errorf("invalid origin")
		// }

		clientId := req.Header.Get(CLIENT_KEY)
		header := http.Header{}
		header.Add(CLIENT_KEY, clientId)
		config.Header = header
		return err
	}
}
func (s *Server) messageHandler() {
	for {
		select {
		case c := <-s.addCh:
			s.clients[c.ID] = c
			log.Printf("Added new client %v \r\n", c.ID)
			log.Printf("Now %v clients connected.", len(s.clients))
		case c := <-s.delCh:
			delete(s.clients, c.ID)
			log.Println("Delete client")
		case err := <-s.errCh:
			log.Println("Error:", err)
		case <-s.doneCh:
			log.Println("Server Done")
			return
		case msg := <-s.messageCh:
			log.Printf("Server received message %v \r\n", msg)
			/* Message Options */

			// Review Commands

			// Generate Party

			// Join Party

			// Leave Party

			// Party Message

			// Game Message

			if msg.Player != nil {
				log.Printf("Searching for party %s \r\n", msg)
				party := s.partySource.get(msg.Player.PartyCode)
				if party != nil {
					log.Printf("Sending message to party %v \r\n", party.Code)
					party.write(msg)
				} else {
					log.Printf("Failed to find party...")
					msg.client.write(simpleMessage{
						client:  msg.client,
						Command: ERROR,
						Detail:  "Party unavailable.",
					})
				}
			} else {
				log.Printf("Not sure what to do with message %v", msg)
				msg.client.write(simpleMessage{
					client:  msg.client,
					Command: ERROR,
					Detail:  "Unhandled message.",
				})
			}
		}
	}
}
func (s *Server) partyGenCodeHandler() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		code := c.Param("code")
		party := s.partySource.host(s, code)
		go party.listen()
		return c.JSON(http.StatusOK, party.Code)
	}
}
func (s *Server) partyCodeHandler() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		code := c.Param("code")
		p := s.partySource.get(code)
		return c.JSON(http.StatusOK, p)
	}
}
func NewServer(origin string) *Server {
	partySource := newPartySource()
	clients := make(map[string]*client)
	addCh := make(chan *client)
	delCh := make(chan *client)
	doneCh := make(chan bool)
	errCh := make(chan error)
	messageCh := make(chan *message)

	return &Server{
		origin,
		partySource,
		clients,
		addCh,
		delCh,
		doneCh,
		errCh,
		messageCh,
	}
}
