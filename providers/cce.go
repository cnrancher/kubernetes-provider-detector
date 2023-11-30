package providers

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const CCE = "cce"

func IsCCE(ctx context.Context, k8sClient kubernetes.Interface) (bool, error) {
	// Look for nodes that have an CCE specific label
	listOpts := metav1.ListOptions{
		LabelSelector: "cce.cloud.com/cce-nodepool",
		// Only need one
		Limit: 1,
	}

	nodes, err := k8sClient.CoreV1().Nodes().List(ctx, listOpts)
	if err != nil {
		return false, err
	}
	if len(nodes.Items) > 0 {
		return true, nil
	}
	return false, nil
}
