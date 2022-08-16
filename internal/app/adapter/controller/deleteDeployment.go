package controller

import (
	"github.com/doyyan/kubernetes/internal/app/application/usecase"
	"github.com/doyyan/kubernetes/internal/app/domain/valueObject"
	"github.com/gin-gonic/gin"
)

func (ctrl Controller) deleteDeployment(c *gin.Context) {
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
	if err := usecase.DeleteDeployment(ctrl.Context, ctrl.Logger, args); err != nil {
		ctrl.processError(c, err)
	}
	c.JSON(200, nil)
}
