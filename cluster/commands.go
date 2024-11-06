package cluster

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(CreateClusterCmd, StatusCmd, DeployCmd)
}

var RootCmd = &cobra.Command{
	Use:   "tufin",
	Short: "Tufin client for managing a Kubernetes cluster",
}

var CreateClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create a k3s Kubernetes cluster",
	Run:   createClusterCmd,
}

var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy MySQL and WordPress pods",
	Run:   deployCmd,
}

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print the status of pods in the default namespace",
	Run:   statusCmd,
}
