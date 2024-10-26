package app

import (
	"catFeeding/app/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func NewService() *Service {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	return &Service{
		DB: cfg.DB,
	}
}

func (s *Service) RegisterRoutes(router *gin.Engine) {
	router.POST("/api/meals", s.createMeal)
	router.GET("/api/meals", s.getMeals)
	router.GET("/api/cats", s.getCats)
	router.GET("/api/users", s.getUsers)
}

func (s *Service) Run() {
	router := gin.Default()
	s.RegisterRoutes(router)
	router.Run()
}

// Handler implementations
func (s *Service) createMeal(ctx *gin.Context) {
	var meal Meal
	if err := ctx.ShouldBindJSON(&meal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := s.DB.Create(&meal)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Load associations
	s.DB.Preload("Cat").Preload("FedBy").First(&meal)

	ctx.JSON(http.StatusCreated, meal)
}

func (s *Service) getMeals(ctx *gin.Context) {
	var meals []Meal
	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")

	query := s.DB.Preload("Cat").Preload("FedBy")

	if startDate != "" && endDate != "" {
		query = query.Where("timestamp BETWEEN ? AND ?", startDate, endDate)
	}

	result := query.Find(&meals)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, meals)
}

func (s *Service) getCats(ctx *gin.Context) {
	var cats []Cat
	result := s.DB.Find(&cats)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, cats)
}

func (s *Service) getUsers(ctx *gin.Context) {
	var users []User
	result := s.DB.Find(&users)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
