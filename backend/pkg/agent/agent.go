package agent

import (
	"context"
	"fmt"

	"github.com/alpineworks/katalog/backend/pkg/agentservice"
	"github.com/michaelpeterswa/go-mtls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type AgentServiceClient struct {
	conn               *grpc.ClientConn
	x509Files          *mtls.X509Files
	agentServiceClient agentservice.AgentServiceClient
}

type AgentServiceClientOption func(*AgentServiceClient)

func NewAgentServiceClient(addr string, options ...AgentServiceClientOption) (*AgentServiceClient, error) {
	var (
		transportCredentials credentials.TransportCredentials
		err                  error
	)

	agentServiceClient := &AgentServiceClient{}

	for _, option := range options {
		option(agentServiceClient)
	}

	if agentServiceClient.x509Files != nil {
		transportCredentials, err = agentServiceClient.x509Files.GenerateTransportCredentials(mtls.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to generate transport credentials: %w", err)
		}
	} else {
		transportCredentials = insecure.NewCredentials()
	}

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &AgentServiceClient{
		conn:               conn,
		agentServiceClient: agentservice.NewAgentServiceClient(conn),
	}, nil
}

func WithMutualTLS(x509Files *mtls.X509Files) AgentServiceClientOption {
	return func(asc *AgentServiceClient) {
		asc.x509Files = x509Files
	}
}

func (asc *AgentServiceClient) Close() error {
	return asc.conn.Close()
}

func (asc *AgentServiceClient) PublishDeployments(ctx context.Context, pdr *agentservice.PublishDeploymentsRequest) (*agentservice.PublishDeploymentsResponse, error) {
	return asc.agentServiceClient.PublishDeployments(ctx, pdr)
}
