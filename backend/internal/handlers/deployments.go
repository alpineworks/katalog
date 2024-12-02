package handlers

import (
	"encoding/json"
	"net/http"

	d "github.com/alpineworks/katalog/backend/internal/deployments"
	"go.uber.org/zap"
)

type DeploymentsHandler struct {
	deploymentsClient *d.DeploymentsClient
	logger            *zap.Logger
}

func NewDeploymentsClient(logger *zap.Logger, deploymentsClient *d.DeploymentsClient) *DeploymentsHandler {
	return &DeploymentsHandler{
		logger:            logger,
		deploymentsClient: deploymentsClient,
	}
}

type GetDeploymentsResponse struct {
	Deployments []d.Deployment `json:"deployments"`
}

func (dh *DeploymentsHandler) GetDeployments(w http.ResponseWriter, r *http.Request) {
	_, span := tracer.Start(r.Context(), "GetDeployments")
	defer span.End()

	deployments, err := dh.deploymentsClient.GetDeployments(r.Context())
	if err != nil {
		dh.logger.Error("error getting deployments", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if deployments == nil {
		deploymentsJSON, err := json.Marshal(GetDeploymentsResponse{Deployments: []d.Deployment{}})
		if err != nil {
			dh.logger.Error("error marshalling deployments", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(deploymentsJSON))
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
