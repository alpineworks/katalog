package grpcserver

import (
	"context"

	"github.com/alpineworks/katalog/backend/internal/deployments"
	"github.com/alpineworks/katalog/backend/pkg/agentservice"
)

type AgentServiceServer struct {
	deploymentsClient *deployments.DeploymentsClient

	agentservice.UnimplementedAgentServiceServer
}

func NewAgentServiceServer(dc *deployments.DeploymentsClient) *AgentServiceServer {
	return &AgentServiceServer{
		deploymentsClient: dc,
	}
}

func (s *AgentServiceServer) PublishDeployments(ctx context.Context, pdr *agentservice.PublishDeploymentsRequest) (*agentservice.PublishDeploymentsResponse, error) {
	for _, deployment := range pdr.Deployments {
		convertedDeployment := deployments.Deployment{
			Name:         deployment.Name,
			Namespace:    deployment.Namespace,
			Cluster:      deployment.Cluster,
			Replicas:     deployment.Replicas,
			TrueReplicas: deployment.TrueReplicas,
			Labels:       deployment.Labels,
		}

		err := s.deploymentsClient.UpsertDeployment(ctx, convertedDeployment)
		if err != nil {
			errMsg := err.Error()
			return &agentservice.PublishDeploymentsResponse{
				Success: false,
				Error:   &errMsg,
			}, err
		}
	}
	return &agentservice.PublishDeploymentsResponse{
		Success: true,
	}, nil
}
