package key

import (
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ClusterNameLabel = "cluster.x-k8s.io/cluster-name"
)

func GetClusterIDFromLabels(t v1.ObjectMeta) string {
	return t.GetLabels()[ClusterNameLabel]
}

func KubeconfigSecretName(clusterName string) string {
	return fmt.Sprintf("%s-kubeconfig", clusterName)
}
