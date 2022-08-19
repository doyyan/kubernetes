package controller

import (
	"net/http"

	"github.com/doyyan/kubernetes/internal/app/application/usecase"
	"github.com/doyyan/kubernetes/internal/app/domain/valueObject"
	"github.com/gin-gonic/gin"
)

//deleteDeployment router to delete a deployment from k8s and DB
func (ctrl Controller) deleteDeployment(c *gin.Context) {
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
	if err := usecase.DeleteDeployment(ctrl.Context, ctrl.Logger, args); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(200, gin.H{
			"message": "deployment deleted",
		})
	}
}
