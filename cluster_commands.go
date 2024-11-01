package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tufin",
	Short: "Tufin client for managing a Kubernetes cluster",
}

var createClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create a k3s Kubernetes cluster",
	Run:   CreateClusterCmd,
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy MySQL and WordPress pods",
	Run:   DeployCmd,
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print the status of pods in the default namespace",
	Run:   StatusCmd,
}
