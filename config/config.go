package config

import (
	"fmt"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"os/user"
	"path/filepath"
)

// GetConfig is a copy of "sigs.k8s.io/controller-runtime"/pkg/client/config.GetConfig without the dependency on other packages of this library
func GetConfig() (config *rest.Config, configErr error) {
	// If the recommended kubeconfig env variable is not specified,
	// try the in-cluster config.
	kubeconfigPath := os.Getenv(clientcmd.RecommendedConfigPathEnvVar)
	if len(kubeconfigPath) == 0 {
		c, err := rest.InClusterConfig()
		if err == nil {
			return c, nil
		}
	}

	// If the recommended kubeconfig env variable is set, or there
	// is no in-cluster config, try the default recommended locations.
	//
	// NOTE: For default config file locations, upstream only checks
	// $HOME for the user's home directory, but we can also try
	// os/user.HomeDir when $HOME is unset.
	//
	// TODO(jlanford): could this be done upstream?

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	if _, ok := os.LookupEnv("HOME"); !ok {
		u, err := user.Current()
		if err != nil {
			return nil, fmt.Errorf("could not get current user: %w", err)
		}
		loadingRules.Precedence = append(loadingRules.Precedence, filepath.Join(u.HomeDir, clientcmd.RecommendedHomeDir, clientcmd.RecommendedFileName))
	}

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{}).ClientConfig()
}

func GetConfigOrDie() *rest.Config {
	cfg, err := GetConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
