package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mca312/jackbox/server/app"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port  int      `yaml:"port"`
	Users []string `yaml:"users"`
}

// Manager for our dependencies
// we'd put in our logging, tracing, token management or other dependencies we'd need across packages
type Manager struct {
	app.UserServiceApi
}

func NewManager() *Manager {
	return &Manager{
		app.NewUserService(),
	}
}

type Server struct {
	*http.Server
}

func main() {
	// load config from yaml file
	cfg := loadConfig()

	// setup server and load test data
	server, err := setupServer(cfg)

	if err != nil {
		fmt.Printf("Error setting up server. %+v", err)
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// properly shutdown server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}

// setupServer manages the routes and init data
func setupServer(cfg *Config) (*Server, error) {
	manager := NewManager()

	// load test data
	for _, u := range cfg.Users {
		configUser := strings.Split(u, ":")
		manager.Register(app.User{Email: configUser[0], Name: configUser[1]})
	}

	router := setupRouter(manager)

	srv := Server{
		&http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Port),
			Handler: router,
		},
	}

	return &srv, nil
}

// setupRouter defines our routes
func setupRouter(manager *Manager) *gin.Engine {
	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // not safe for production
	r.Use(cors.New(config))
	r.OPTIONS("/*cors", func(c *gin.Context) {})

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"health": "check"})
	})

	r.POST("/login", manager.LoginHandler)
	r.POST("/register", manager.RegisterHandler)
	r.GET("/user/:user", manager.UserHandler)
	return r
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

	return &config
}
