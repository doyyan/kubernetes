package controller

import (
	"net/http"

	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql/model"
	"github.com/doyyan/kubernetes/internal/app/application/usecase"
	"github.com/doyyan/kubernetes/internal/app/domain/valueObject"
	"github.com/gin-gonic/gin"
)

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
		DeploymentRepository: deploymentRepository,
	}
	if err := usecase.CreateDeployment(ctrl.Context, ctrl.Logger, args); err != nil {
		ctrl.processError(c, err)
	}
	c.JSON(200, deployment)
}

func (ctrl Controller) processError(c *gin.Context, err error) {
	ctrl.Logger.Error(err)
	c.AbortWithError(http.StatusBadRequest, err)
}
