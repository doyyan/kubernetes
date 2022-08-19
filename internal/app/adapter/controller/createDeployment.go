package controller

import (
	"net/http"

	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql/model"
	"github.com/doyyan/kubernetes/internal/app/application/usecase"
	"github.com/doyyan/kubernetes/internal/app/domain/valueObject"
	"github.com/gin-gonic/gin"
)

//createDeployment calls functions to create k8s deployment and store to DB
func (ctrl Controller) createDeployment(c *gin.Context) {
	deployment := model.Deployment{}
	if err := c.BindJSON(&deployment); err != nil {
		ctrl.processError(c, err)
		return
	}
	dep := valueObject.Deployment{
		Namespace:     deployment.NameSpace,
		Name:          deployment.Name,
		ContainerPort: deployment.ContainerPort,
		ContainerName: deployment.ContainerName,
		Image:         deployment.Image,
		Labels:        deployment.Labels,
		Replicas:      deployment.Replicas,
	}
	args := usecase.CreateDeploymentArgs{
		Deployment:           dep,
		DeploymentRepository: ctrl.deploymentRepository,
	}
	if err := usecase.CreateDeployment(ctrl.Context, ctrl.Logger, args); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(200, gin.H{
			"message": "deployment created",
		})
	}
}

func (ctrl Controller) processError(c *gin.Context, err error) {
	ctrl.Logger.Error(err)
	c.AbortWithError(http.StatusBadRequest, err)
}
