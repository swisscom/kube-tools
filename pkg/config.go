package kubetools

import (
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func GetConfig() (*api.Config, error) {
	clientConfigLoadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	return clientConfigLoadingRules.Load()
}
