package appointment

import (
	"github.com/test/usecases"
	"github.com/gin-gonic/gin"
)

type appointmentHandler struct {
	appointmentUseCase usecases.AppointmentUseCase
}

func NewEndpointHttpHandler(ginEng *gin.Engine, uc usecases.AppointmentUseCase) {
	handler := appointmentHandler{appointmentUseCase: uc}
	r := ginEng.Group("apis")
	r.GET("/appointment", handler.List)
	r.POST("/appointment", handler.Create)
	r.DELETE("/appointment/:id", handler.Delete)
}
