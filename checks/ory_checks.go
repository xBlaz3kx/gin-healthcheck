package checks

import (
	"context"
	"time"

	ketoClient "github.com/ory/keto-client-go"
	kratosClient "github.com/ory/kratos-client-go"
)

type KetoCheck struct {
	client  *ketoClient.APIClient
	timeout int
}

func NewKetoCheck(timeout int, client *ketoClient.APIClient) *KetoCheck {
	return &KetoCheck{
		timeout: timeout,
		client:  client,
	}
}

func (k *KetoCheck) Pass() bool {
	if k.client == nil {
		return false
	}

	timeout := time.Second * time.Duration(k.timeout)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, h, err := k.client.MetadataApi.IsAlive(ctx).Execute()
	if err != nil {
		return false
	}

	return h.StatusCode == 200
}

func (k *KetoCheck) Name() string {
	return "keto"
}

type KratosCheck struct {
	client  *kratosClient.APIClient
	timeout int
}

func NewKratosCheck(timeout int, client *kratosClient.APIClient) *KratosCheck {
	return &KratosCheck{
		client:  client,
		timeout: timeout,
	}
}

func (k *KratosCheck) Pass() bool {
	if k.client == nil {
		return false
	}

	timeout := time.Second * time.Duration(k.timeout)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, h, err := k.client.MetadataApi.IsAlive(ctx).Execute()
	if err != nil {
		return false
	}

	return h.StatusCode == 200
}

func (k *KratosCheck) Name() string {
	return "kratos"
}
