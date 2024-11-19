package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alpineworks/katalog/backend/internal/deployments"
	"go.uber.org/zap"
)

type DeploymentsHandler struct {
	deploymentsClient *deployments.DeploymentsClient
	logger            *zap.Logger
}

func NewDeploymentsClient(logger *zap.Logger, deploymentsClient *deployments.DeploymentsClient) *DeploymentsHandler {
	return &DeploymentsHandler{
		logger:            logger,
		deploymentsClient: deploymentsClient,
	}
}

type GetDeploymentsResponse struct {
	Deployments []deployments.Deployment `json:"deployments"`
}

func (dh *DeploymentsHandler) GetDeployments(w http.ResponseWriter, r *http.Request) {
	deployments, err := dh.deploymentsClient.GetDeployments(r.Context())
	if err != nil {
		dh.logger.Error("error getting deployments", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if deployments == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	deploymentsJSON, err := json.Marshal(GetDeploymentsResponse{Deployments: deployments})
	if err != nil {
		dh.logger.Error("error marshalling deployments", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(deploymentsJSON))
}
