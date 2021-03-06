// Copyright 2020 gRPC authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package kubehelpers contains helper functions for interacting with Kubernetes
// clusters and resources.
package kubehelpers

import (
	"k8s.io/client-go/kubernetes"

	// This side-effect import is required by GKE.
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// ConnectWithinCluster uses Kubernetes utility functions to locate
// the credentials within the cluster and connect to the API.
func ConnectWithinCluster() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

// ConnectWithConfig takes an absolute path to a kube config file
// which is used to connect to the Kubernetes API.
func ConnectWithConfig(abspath string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", abspath)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}
