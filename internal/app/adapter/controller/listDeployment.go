package controller

import (
	"github.com/doyyan/kubernetes/internal/app/application/usecase"
	"github.com/gin-gonic/gin"
)

func (ctrl Controller) listDeployment(c *gin.Context) {
	args := usecase.CreateDeploymentArgs{
		DeploymentRepository: ctrl.deploymentRepository,
	}
	deployment, err := usecase.ListDeployment(ctrl.Context, ctrl.Logger, args)
	if err != nil {
		ctrl.processError(c, err)
	}
	if len(deployment) == 0 {
		c.JSON(404, gin.H{
			"message": "no deployments found",
		})
	} else {
		c.JSON(200, deployment)
	}
}
