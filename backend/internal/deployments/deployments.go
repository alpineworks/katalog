package deployments

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"time"

	"fmt"

	"github.com/alpineworks/katalog/backend/internal/dragonfly"
	"github.com/alpineworks/katalog/backend/internal/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	uuid "github.com/satori/go.uuid"
)

type DeploymentsClient struct {
	postgresClient  *postgres.PostgresClient
	dragonflyClient *dragonfly.DragonflyClient
}

func NewDeploymentsClient(postgresClient *postgres.PostgresClient, dragonflyClient *dragonfly.DragonflyClient) *DeploymentsClient {
	return &DeploymentsClient{
		postgresClient:  postgresClient,
		dragonflyClient: dragonflyClient,
	}
}

func (dc *DeploymentsClient) Close() {
	// noop
}

type Deployment struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Namespace    string            `json:"namespace"`
	Cluster      string            `json:"cluster"`
	Replicas     int32             `json:"replicas"`
	TrueReplicas int32             `json:"true_replicas"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	Labels       map[string]string `json:"labels"`
}

//go:embed queries/get_deployments.pgsql
var getDeploymentsSQL string

func (dc *DeploymentsClient) GetDeployments(ctx context.Context) ([]Deployment, error) {
	deploymentsJSON, err := dc.dragonflyClient.GetClient().Get(ctx, "current-deployments").Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("failed to get deployments from dragonfly: %w", err)
	} else if deploymentsJSON != "" {
		var deployments []Deployment
		err := json.Unmarshal([]byte(deploymentsJSON), &deployments)

		// TODO: should just log and continue here
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal deployments: %w", err)
		}

		return deployments, nil
	}

	res, err := dc.postgresClient.Client.Query(ctx, getDeploymentsSQL)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("failed to query deployments: %w", err)
	} else if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	defer res.Close()

	var deployments []Deployment
	for res.Next() {
		var deployment Deployment
		err := res.Scan(&deployment.ID, &deployment.Name, &deployment.Namespace, &deployment.Cluster, &deployment.Replicas, &deployment.TrueReplicas, &deployment.CreatedAt, &deployment.UpdatedAt, &deployment.Labels)
		if err != nil {
			return nil, fmt.Errorf("failed to scan deployment: %w", err)
		}

		deployments = append(deployments, deployment)
	}

	deploymentsBytes, err := json.Marshal(deployments)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal deployments: %w", err)
	}

	err = dc.dragonflyClient.GetClient().Set(ctx, "current-deployments", string(deploymentsBytes), 5*time.Minute).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to set deployments in dragonfly: %w", err)
	}

	return deployments, nil
}

//go:embed queries/get_deployment_by_name_namespace_cluster.pgsql
var getDeploymentByNameNamespaceClusterSQL string

func (dc *DeploymentsClient) GetDeploymentIdByNameNamespaceCluster(ctx context.Context, name, namespace, cluster string) (string, error) {
	var ID string
	err := dc.postgresClient.Client.QueryRow(ctx, getDeploymentByNameNamespaceClusterSQL, name, namespace, cluster).Scan(&ID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return "", fmt.Errorf("failed to query deployment: %w", err)
	} else if errors.Is(err, pgx.ErrNoRows) {
		return "", nil
	}

	return ID, nil
}

//go:embed queries/upsert_deployment.pgsql
var upsertDeploymentSQL string

func (dc *DeploymentsClient) UpsertDeployment(ctx context.Context, deployment Deployment) error {
	id, err := dc.GetDeploymentIdByNameNamespaceCluster(ctx, deployment.Name, deployment.Namespace, deployment.Cluster)
	if err != nil {
		return fmt.Errorf("failed to get deployment id: %w", err)
	}

	if id == "" {
		id = uuid.NewV4().String()
	}

	_, err = dc.postgresClient.Client.Exec(ctx, upsertDeploymentSQL, id, deployment.Name, deployment.Namespace, deployment.Cluster, deployment.Replicas, deployment.TrueReplicas, deployment.Labels, time.Now())
	if err != nil {
		return fmt.Errorf("failed to upsert deployment: %w", err)
	}

	return nil
}
