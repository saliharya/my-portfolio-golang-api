package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portfolio-arya-service/internal/domain"
)

type AboutHandler struct {
	aboutUsecase domain.AboutUsecase
}

func NewAboutHandler(r *gin.Engine, aboutUsecase domain.AboutUsecase) {
	handler := &AboutHandler{
		aboutUsecase: aboutUsecase,
	}

	about := r.Group("/about")
	{
		about.GET("", handler.GetAbout)
		about.PATCH("", handler.UpdateAbout)
	}
}

func (h *AboutHandler) GetAbout(c *gin.Context) {
	about, err := h.aboutUsecase.GetAbout()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, about)
}

func (h *AboutHandler) UpdateAbout(c *gin.Context) {
	var about domain.About
	if err := c.ShouldBindJSON(&about); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.aboutUsecase.UpdateAbout(about)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}
