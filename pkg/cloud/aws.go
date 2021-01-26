package cloud

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetAWSResources() OwnedResources {
	return OwnedResources{
		newTestResource(),
	}
}

func newTestResource() *v1.Service {
	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{{
				Protocol: "TCP",
				Port:     8080,
			}},
		},
	}
	svc.APIVersion = "v1"
	svc.Kind = "Service"
	return svc
}
