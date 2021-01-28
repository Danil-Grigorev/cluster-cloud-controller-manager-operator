package platform

import (
	"context"

	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PlatformOwner interface {
	List(ctx context.Context, c client.Client) ([]client.ObjectKey, error)
	GetObjects(ctx context.Context, с client.Client, key client.ObjectKey) (objects []client.Object, error error)
	Object() client.Object
}

func GetPlatform(platformKey string, namespace string) PlatformOwner {
	switch platformKey {
	case ClusterKey:
		return &ClusterOwner{namespace}
	case InfrastructureKey:
		return &InfrastrucutreOwner{}
	default:
		klog.Infof("Unknown platform: defaulting to '%T'", InfrastrucutreOwner{})
		return &InfrastrucutreOwner{}
	}
}
