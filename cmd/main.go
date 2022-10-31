package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ginmills/cheers"
	"github.com/ginmills/ginmill"
)

type Server struct {
	ginmill.Server
	cheers.ICheers
}

func (s *Server) start() {
	s.Engine = gin.New()

	s.Engine.Use(
		gin.Recovery(),
	)

	s.With(cheers.Features(s))
	server := http.Server{
		//		TLSConfig: tlsConfig,
		Handler: s.Engine,
		Addr:    ":8080",
	}

	server.ListenAndServe()
}

// response to cheers
func (s *Server) Cheers(c *gin.Context) {
	c.JSON(http.StatusOK, "Cheers")
}

func main() {
	s := new(Server)
	s.start()
}
