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

func Test_RolloutStatusDeployment(t *testing.T) {
	tests := map[string]struct {
		err            error
		testDep        model.Deployment
		testcase       domainrepo.IDeploymentRepo
		output         string
		httpReturnCode int
	}{
		"pass getRolloutStatusSuccess": {
			err:            nil,
			testDep:        DeploymentData(),
			testcase:       getRolloutStatusSuccess(),
			output:         "rollout in progress",
			httpReturnCode: 200,
		},
		"pass getRolloutStatusFail": {
			err:            errors.New(" DB save failure"),
			testcase:       getRolloutStatusFail(),
			testDep:        DeploymentData(),
			httpReturnCode: 500,
			output:         "",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			router := setupGetRolloutStatusRouter(test.testcase)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/deployment/status", nil)
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

func setupGetRolloutStatusRouter(testcase domainrepo.IDeploymentRepo) *gin.Engine {
	r := gin.Default()
	cont := CreateController()
	cont.deploymentRepository = testcase
	r.GET("/deployment/status", cont.getRolloutStatus)
	return r
}

func getRolloutStatusSuccess() domainrepo.IDeploymentRepo {
	depRepoMock := domainrepotest.IDeploymentRepoMock{GetRolloutStatusFunc: func(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) (string, bool, error) {
		return "rollout in progress", false, nil
	},
		GetFunc: func(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) (domain.Deployment, error) {
			return domain.Deployment{Name: "testDeployment"}, nil
		},
	}
	return &depRepoMock
}

func getRolloutStatusFail() domainrepo.IDeploymentRepo {
	depRepoMock := domainrepotest.IDeploymentRepoMock{GetRolloutStatusFunc: func(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) (string, bool, error) {
		return "", false, errors.New(" DB save failure")
	},
		GetFunc: func(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) (domain.Deployment, error) {
			return domain.Deployment{Name: "testDeployment"}, nil
		}}
	return &depRepoMock
}
