package controller

import (
	"net/http"

	"github.com/doyyan/kubernetes/internal/app/application/usecase"
	"github.com/doyyan/kubernetes/internal/app/domain/valueObject"
	"github.com/gin-gonic/gin"
)

//getDeployment router to get a single deployment from the DB
func (ctrl Controller) getDeployment(c *gin.Context) {
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
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		if deployment.Name == "" {
			c.JSON(404, gin.H{
				"message": "deployment not found",
			})
		} else {
			c.JSON(200, deployment)
		}
	}

}
