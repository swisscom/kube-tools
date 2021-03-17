package main

import (
	"fmt"
	kubetools "github.com/swisscom/kube-tools/pkg"
	"os"
)

/*
	Super simple app that check your KUBECONFIG and returns the current context
 */
func main(){
	cfg, err := kubetools.GetConfig()
	if err != nil {
		kubetools.Errorf("unable to get kubernetes config: %v", err)
		os.Exit(1)
	}

	ctx, ok := cfg.Contexts[cfg.CurrentContext]
	if !ok {
		kubetools.Errorf("unable to get current context")
		os.Exit(2)
	}

	fmt.Printf("%s@%s:%s",
		ctx.AuthInfo,
		ctx.Cluster,
		ctx.Namespace)
}