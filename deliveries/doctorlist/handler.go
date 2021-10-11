package doctorlist

import (
	"github.com/test/usecases"
	"github.com/gin-gonic/gin"
)

type doctorHandler struct {
	doctorUseCase usecases.DoctorUseCase
}

func NewEndpointHttpHandler(ginEng *gin.Engine, uc usecases.DoctorUseCase) {
	handler := doctorHandler{doctorUseCase: uc}
	r := ginEng.Group("admin")
	r.GET("/doctors", handler.List)
}
