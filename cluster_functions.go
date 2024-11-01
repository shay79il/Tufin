package main

import (
	"fmt"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"os/exec"
	"strings"
)

func CreateClusterCmd(cmd *cobra.Command, args []string) {

	clusterName := "tufin-cluster"
	k3dCmd := exec.Command("k3d", "cluster", "get", clusterName)
	if err := k3dCmd.Run(); err == nil {
		fmt.Printf("Cluster '%s' already exists.\n", clusterName)
		return
	} else {
		k3dCmd = exec.Command("k3d", "cluster", "create", "--api-port", "6550", "-p", "8080:80@loadbalancer", clusterName)
		k3dCmd.Stdout = os.Stdout
		k3dCmd.Stderr = os.Stderr
		if err := k3dCmd.Run(); err != nil {
			fmt.Println("Error creating k3d cluster:", err)
			return
		}
	}

	fmt.Println("k3s cluster deployed successfully using k3d!")
}

func StatusCmd(cmd *cobra.Command, args []string) {

	pods, err := k8sClient.CoreV1().Pods("default").List(cmd.Context(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error getting pods:", err)
		return
	}

	fmt.Println("Pod Name\t\tStatus")
	for _, pod := range pods.Items {
		status := strings.ToUpper(string(pod.Status.Phase))
		fmt.Printf("%s\t\t%s\n", pod.Name, status)
	}
}

func DeployCmd(cmd *cobra.Command, args []string) {

	_, err := k8sClient.CoreV1().Pods("default").Get(cmd.Context(), "mysql", metav1.GetOptions{})
	if err == nil {
		fmt.Println("MySQL pod already exists.")
	} else {
		_, err := k8sClient.CoreV1().Pods("default").Create(cmd.Context(), mysqlPod, metav1.CreateOptions{})
		if err != nil {
			fmt.Println("Error creating MySQL pod:", err)
			return
		}
	}

	_, err = k8sClient.CoreV1().Services("default").Get(cmd.Context(), "mysql", metav1.GetOptions{})
	if err == nil {
		fmt.Println("MySQL service already exists.")
	} else {
		_, err = k8sClient.CoreV1().Services("default").Create(cmd.Context(), mysqlService, metav1.CreateOptions{})
		if err != nil {
			fmt.Println("Error creating MySQL service:", err)
			return
		}
	}
	_, err = k8sClient.CoreV1().Pods("default").Get(cmd.Context(), "wordpress", metav1.GetOptions{})
	if err == nil {
		fmt.Println("Wordpress pod already exists.")
	} else {
		_, err = k8sClient.CoreV1().Pods("default").Create(cmd.Context(), wordpressPod, metav1.CreateOptions{})
		if err != nil {
			fmt.Println("Error creating WordPress pod:", err)
			return
		}
	}

	_, err = k8sClient.CoreV1().Services("default").Get(cmd.Context(), "wordpress", metav1.GetOptions{})
	if err == nil {
		fmt.Println("Wordpress service already exists.")
	} else {
		_, err = k8sClient.CoreV1().Services("default").Create(cmd.Context(), wordpressService, metav1.CreateOptions{})
		if err != nil {
			fmt.Println("Error creating MySQL service:", err)
			return
		}
	}

	fmt.Println("MySQL and WordPress pods deployed successfully!")
}
