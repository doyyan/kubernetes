package controller

import (
	"github.com/doyyan/kubernetes/internal/app/adapter/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

var (
	deploymentRepository = repository.Deployment{}
)

type Controller struct {
	Context context.Context
	Logger  *logrus.Logger
}

func (ctrl Controller) Router() *gin.Engine {
	r := gin.Default()
	r.POST("/deployment", ctrl.createDeployment)
	return r
}
