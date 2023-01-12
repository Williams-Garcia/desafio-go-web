package routes

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/handlers"
	"github.com/bootcamp-go/desafio-go-web/cmd/handlers/ping"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Router struct {
	en *gin.Engine
	db []domain.Ticket
}

func NewRouter(en *gin.Engine, db []domain.Ticket) *Router {
	return &Router{en: en, db: db}
}

func (r *Router) SetRoutes() {
	r.SetPing()
	r.SetTickets()
}

// website
func (r *Router) SetTickets() {
	repo := tickets.NewRepository(r.db)
	service := tickets.NewService(&repo)
	handler := handlers.NewHandler(service)

	routerTickets := r.en.Group("ticket")
	routerTickets.GET("getByCountry/:dest", handler.GetTicketsByCountry())
	routerTickets.GET("getAverage/:dest", handler.AverageDestination())
}

func (r *Router) SetPing() {
	r.en.GET("/ping", ping.Ping)
}
