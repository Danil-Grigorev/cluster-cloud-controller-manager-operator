package platform

import (
	"context"

	configv1 "github.com/openshift/api/config/v1"

	"github.com/openshift/cluster-cloud-controller-manager-operator/pkg/cloud"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"k8s.io/klog/v2"
)

var _ PlatformOwner = &InfrastrucutreOwner{}

type InfrastrucutreOwner struct {
	objects []configv1.Infrastructure
}

func (o *InfrastrucutreOwner) Object() client.Object {
	return &configv1.Infrastructure{}
}

func (o *InfrastrucutreOwner) List(ctx context.Context, c client.Client) ([]client.ObjectKey, error) {
	infraList := &configv1.InfrastructureList{}
	if err := c.List(ctx, infraList); err != nil {
		klog.Errorf("Unable to retrive list of platform %T objects: %v", infraList, err)
		return nil, err
	}

	keys := []client.ObjectKey{}
	for _, infra := range infraList.Items {
		keys = append(keys, client.ObjectKeyFromObject(&infra))
	}
	return keys, nil
}

func (o *InfrastrucutreOwner) GetObjects(ctx context.Context, с client.Client, key client.ObjectKey) ([]client.Object, error) {
	infra := &configv1.Infrastructure{}
	err := с.Get(ctx, key, infra)
	if err != nil {
		klog.Errorf("Unable to retrive platform %T object: %v", infra, err)
		return nil, err
	}

	return getResources(infra.Status.Platform), nil
}

func getResources(platformType configv1.PlatformType) []client.Object {
	switch platformType {
	case configv1.AWSPlatformType:
		return cloud.GetAWSResources()
	default:
		klog.Warning("No recognized cloud provider platform found in infrastructure")
	}
	return nil
}
