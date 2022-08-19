package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql/model"
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/doyyan/kubernetes/internal/app/domain/domainrepo"
	domainrepotest "github.com/doyyan/kubernetes/internal/app/domain/domainrepo/mocks"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func Test_ListDeployment(t *testing.T) {
	tests := map[string]struct {
		err            error
		testDep        model.Deployment
		testcase       domainrepo.IDeploymentRepo
		output         string
		httpReturnCode int
	}{
		"pass listDeploymentSuccess": {
			err:            nil,
			testDep:        DeploymentData(),
			testcase:       listDeploymentSuccess(),
			output:         "testDeployment",
			httpReturnCode: 200,
		},
		"pass listDeploymentFail": {
			err:            errors.New(" DB save failure"),
			testcase:       listDeploymentFail(),
			testDep:        DeploymentData(),
			httpReturnCode: 500,
			output:         "",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			router := setupListDeploymentRouter(test.testcase)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/deployment/all", nil)
			router.ServeHTTP(w, req)
			assert.Equal(t, test.httpReturnCode, w.Code)
			if w.Code == 200 {
				var output []domain.Deployment
				err := json.Unmarshal(w.Body.Bytes(), &output)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, test.output, output[0].Name)
			}
		})
	}
}

func setupListDeploymentRouter(testcase domainrepo.IDeploymentRepo) *gin.Engine {
	r := gin.Default()
	cont := CreateController()
	cont.deploymentRepository = testcase
	r.GET("/deployment/all", cont.listDeployment)
	return r
}

func listDeploymentSuccess() domainrepo.IDeploymentRepo {
	depRepoMock := domainrepotest.IDeploymentRepoMock{ListFunc: func(ctx context.Context, logger *logrus.Logger) ([]domain.Deployment, error) {
		return []domain.Deployment{{Name: "testDeployment"}}, nil
	}}
	return &depRepoMock
}

func listDeploymentFail() domainrepo.IDeploymentRepo {
	depRepoMock := domainrepotest.IDeploymentRepoMock{ListFunc: func(ctx context.Context, logger *logrus.Logger) ([]domain.Deployment, error) {
		return nil, errors.New(" DB save failure")
	}}
	return &depRepoMock
}
