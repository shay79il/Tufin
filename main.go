package main

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

var k8sClient *kubernetes.Clientset

func init() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		fmt.Println("Error getting Kubernetes config:", err)
		os.Exit(1)
	}

	k8sClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Error creating Kubernetes clientset:", err)
		os.Exit(1)
	}
}

func main() {
	rootCmd.AddCommand(createClusterCmd, deployCmd, statusCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
