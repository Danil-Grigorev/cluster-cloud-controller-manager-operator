package platform

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PlatformOwner interface {
	List(ctx context.Context, c client.Client) ([]client.ObjectKey, error)
	GetObjects(ctx context.Context, с client.Client, key client.ObjectKey) (objects []client.Object, error error)
	Object() client.Object
}
