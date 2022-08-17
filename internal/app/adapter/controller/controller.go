package controller

import (
	"github.com/doyyan/kubernetes/internal/app/adapter/kubernetes"
	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql"
	"github.com/doyyan/kubernetes/internal/app/adapter/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type Controller struct {
	Context              context.Context
	Logger               *logrus.Logger
	deploymentRepository repository.Deployment
}

func (ctrl Controller) Router() *gin.Engine {
	ctrl.deploymentRepository = repository.Deployment{postgresql.GetDB(), kubernetes.Kube{}}
	r := gin.Default()
	r.POST("/deployment", ctrl.createDeployment)
	r.GET("/deployment", ctrl.getDeployment)
	r.DELETE("/deployment", ctrl.deleteDeployment)
	r.GET("/deployment/status", ctrl.getRolloutStatus)
	return r
}
