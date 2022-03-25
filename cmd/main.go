package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/mca312/jackbox/server/app"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port  int      `yaml:"port"`
	Users []string `yaml:"users"`
}

type Server struct {
	*http.Server
}

func main() {
	cfg := loadConfig()

	server, err := setupServer(cfg)
	if err != nil {
		fmt.Printf("Error setting up server. %+v", err)
	}

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}
}

func setupServer(cfg *Config) (*Server, error) {
	router := gin.New()

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"health": "check"})
	})

	router.POST("/login", LoginHandler)
	router.POST("/register", RegisterHandler)
	router.GET("/user/:user", UserHandler)

	srv := Server{
		&http.Server{
			Addr:    fmt.Sprintf("localhost:%d", cfg.Port),
			Handler: router,
		},
	}

	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-done
		srv.Close()
	}()

	return &srv, nil
}

func loadConfig() *Config {
	f, err := ioutil.ReadFile("../config.yml")
	if err != nil {
		panic(err)
	}

	var config Config
	if err = yaml.Unmarshal(f, &config); err != nil {
		panic(err)
	}

	for _, u := range config.Users {
		app.Register(app.User{Email: u})
	}

	return &config
}
