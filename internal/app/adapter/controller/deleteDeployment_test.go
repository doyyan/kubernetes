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

func Test_DeleteDeployment(t *testing.T) {
	tests := map[string]struct {
		err            error
		testDep        model.Deployment
		testcase       domainrepo.IDeploymentRepo
		output         string
		httpReturnCode int
	}{
		"pass createDeploymentSuccess": {
			err:            nil,
			testDep:        DeploymentData(),
			testcase:       deleteDeploymentSuccess(),
			output:         "deployment deleted",
			httpReturnCode: 200,
		},
		"pass createDeploymentFail": {
			err:            errors.New(" DB save failure"),
			testcase:       deleteDeploymentFail(),
			testDep:        DeploymentData(),
			httpReturnCode: 500,
			output:         "",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			router := setupDeletionRouter(test.testcase)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, "/deployment", nil)
			q := req.URL.Query()            // Get a copy of the query values.
			q.Add("name", "testDeployment") // Add a new value to the set.
			req.URL.RawQuery = q.Encode()   // Encode and assign back to the original query.
			router.ServeHTTP(w, req)
			assert.Equal(t, test.httpReturnCode, w.Code)
			if w.Code == 200 {
				var output map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &output)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, test.output, output["message"])
			}
		})
	}
}

func setupDeletionRouter(testcase domainrepo.IDeploymentRepo) *gin.Engine {
	r := gin.Default()
	cont := CreateController()
	cont.deploymentRepository = testcase
	r.DELETE("/deployment", cont.deleteDeployment)
	return r
}

func deleteDeploymentSuccess() domainrepo.IDeploymentRepo {
	depRepoMock := domainrepotest.IDeploymentRepoMock{DeleteFunc: func(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) error {
		return nil
	}}
	return &depRepoMock
}

func deleteDeploymentFail() domainrepo.IDeploymentRepo {
	depRepoMock := domainrepotest.IDeploymentRepoMock{DeleteFunc: func(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) error {
		return errors.New(" DB save failure")
	}}
	return &depRepoMock
}
