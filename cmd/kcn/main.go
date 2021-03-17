package main

import (
	"fmt"
	kubetools "github.com/swisscom/kube-tools/pkg"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func usage(){
	fmt.Printf("Usage: %s NEW-NAMESPACE\n", os.Args[0])
}

func main(){
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	wantedNamespace := os.Args[1]

	clientConfigLoadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configPath := clientConfigLoadingRules.GetDefaultFilename()

	if configPath == "" {
		kubetools.Errorf("unable to determine Kubernetes config path: %v")
		os.Exit(2)
	}

	config, err := clientConfigLoadingRules.Load()

	if err != nil {
		kubetools.Errorf("unable to load config: %v", err)
		os.Exit(3)
	}

	ctx, ok := config.Contexts[config.CurrentContext]
	if !ok {
		kubetools.Errorf("unable to get current context")
		os.Exit(4)
	}

	clientConfigLoadingRules.GetDefaultFilename()

	ctx.Namespace = wantedNamespace
	err = clientcmd.WriteToFile(*config, configPath)

	if err != nil {
		kubetools.Errorf("unable to write config to file: %v", err)
		os.Exit(5)
	}
}