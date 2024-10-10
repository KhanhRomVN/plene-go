package routes

import (
	"pleno-go/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(r *gin.Engine, db *pgxpool.Pool) {
	r.GET("/", handlers.HelloHandler())
	r.POST("/register", handlers.Register(db))
	r.POST("/login", handlers.Login(db, "your-jwt-secret"))
}
