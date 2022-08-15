package controller

import (
	"github.com/doyyan/kubernetes/internal/app/adapter/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	deploymentRepository = repository.Deployment{}
)

type Controller struct {
	Logger *logrus.Logger
}

func (ctrl Controller) Router() *gin.Engine {
	r := gin.Default()
	r.POST("/deployment", ctrl.createDeployment)
	return r
}
