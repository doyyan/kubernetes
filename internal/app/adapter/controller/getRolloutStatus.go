package controller

import (
	"github.com/doyyan/kubernetes/internal/app/application/usecase"
	"github.com/doyyan/kubernetes/internal/app/domain/valueObject"
	"github.com/gin-gonic/gin"
)

func (ctrl Controller) getRolloutStatus(c *gin.Context) {
	namespace := c.DefaultQuery("namespace", "default")
	name := c.Query("name")
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
	if deployment.Name == "" {
		c.JSON(404, gin.H{
			"message": "deployment not found",
		})
		return
	}
	message, status, err := usecase.GetRolloutStatus(ctrl.Context, ctrl.Logger, args)
	if err != nil {
		ctrl.processError(c, err)
	}
	var rolloutStatus string
	switch status {
	case true:
		rolloutStatus = "Process Complete"
	default:
		rolloutStatus = "Process in Progress"

	}
	c.JSON(200, gin.H{
		"rollout Status": rolloutStatus,
		"message":        message,
	})

}
