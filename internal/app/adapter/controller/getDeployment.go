package controller

import (
	"github.com/doyyan/kubernetes/internal/app/application/usecase"
	"github.com/doyyan/kubernetes/internal/app/domain/valueObject"
	"github.com/gin-gonic/gin"
)

func (ctrl Controller) getDeployment(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")
	dep := valueObject.Deployment{
		Namespace: namespace,
		Name:      name,
	}
	args := usecase.CreateDeploymentArgs{
		Deployment:           dep,
		DeploymentRepository: ctrl.deploymentRepository,
	}
	deployment, err := usecase.GetDeployment(ctrl.Context, ctrl.Logger, args)
	if err != nil {
		ctrl.processError(c, err)
	}
	c.JSON(200, deployment)
}
