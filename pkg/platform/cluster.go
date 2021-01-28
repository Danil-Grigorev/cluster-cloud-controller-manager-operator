package platform

import (
	"context"

	"github.com/openshift/cluster-cloud-controller-manager-operator/pkg/cloud"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
	clusterctlv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ PlatformOwner = &ClusterOwner{}

const ClusterKey = "cluster"

type ClusterOwner struct {
	namespace string
}

func (o *ClusterOwner) Object() client.Object {
	return &clusterctlv1.Cluster{}
}

func (o *ClusterOwner) List(ctx context.Context, c client.Client) ([]client.ObjectKey, error) {
	clusterList := &clusterctlv1.ClusterList{}
	if err := c.List(ctx, clusterList, client.InNamespace(o.namespace)); err != nil {
		klog.Errorf("Unable to retrive list of platform %T objects: %v", clusterList, err)
		return nil, err
	}

	keys := []client.ObjectKey{}
	for _, cluster := range clusterList.Items {
		keys = append(keys, client.ObjectKeyFromObject(&cluster))
	}
	return keys, nil
}

func (o *ClusterOwner) GetObjects(ctx context.Context, с client.Client, key client.ObjectKey) ([]client.Object, error) {
	clsuter := &clusterctlv1.Cluster{}
	err := с.Get(ctx, key, clsuter)
	if err != nil {
		klog.Errorf("Unable to retrive platform %T object: %v", clsuter, err)
		return nil, err
	}

	return getResourcesForCluster(clsuter.Spec.InfrastructureRef), nil
}

func getResourcesForCluster(platformType *corev1.ObjectReference) []client.Object {
	switch platformType.Kind {
	case "AWSCluster":
		return cloud.GetAWSResources()
	default:
		klog.Warning("No recognized cloud provider platform found in infrastructure")
	}
	return nil
}
