package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portfolio-arya-service/internal/domain"
	"strconv"
)

type ServicesHandler struct {
	servicesUsecase domain.ServicesUsecase
}

func NewServicesHandler(r *gin.Engine, servicesUsecase domain.ServicesUsecase) {
	handler := &ServicesHandler{
		servicesUsecase: servicesUsecase,
	}

	services := r.Group("/services")
	{
		services.GET("", handler.GetAllServices)
		services.POST("", handler.CreateService)
		services.PUT("/:id", handler.UpdateService)
		services.DELETE("/:id", handler.DeleteService)
		services.PATCH("/summary", handler.UpdateServicesSummary)
	}
}

func (h *ServicesHandler) GetAllServices(c *gin.Context) {
	services, err := h.servicesUsecase.GetServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

func (h *ServicesHandler) CreateService(c *gin.Context) {
	var newItemService domain.ServiceItem
	if err := c.ShouldBindJSON(&newItemService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.servicesUsecase.CreateService(newItemService)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, created)
}

func (h *ServicesHandler) UpdateService(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service id: " + idParam})
		return
	}

	var updatedItemService domain.ServiceItem
	if err := c.ShouldBindJSON(&updatedItemService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedItemService.Id = id
	updated, err := h.servicesUsecase.UpdateService(updatedItemService)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *ServicesHandler) DeleteService(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service id: " + idParam})
		return
	}

	if err := h.servicesUsecase.DeleteService(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted"})
}

func (h *ServicesHandler) UpdateServicesSummary(c *gin.Context) {
	var payload struct {
		Summary string `json:"summary"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.servicesUsecase.UpdateSummary(payload.Summary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}
