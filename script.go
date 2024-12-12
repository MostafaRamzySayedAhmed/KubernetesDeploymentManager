package main

import (
	"fmt"
	"os/exec"
	"log"
)

func main() {
	// Apply Kubernetes deployment YAML file
	cmd := exec.Command("kubectl", "apply", "-f", "deployment.yaml")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deployment applied successfully.")

	// Expose the deployment as a service
	cmd = exec.Command("kubectl", "expose", "deployment", "my-app", "--type=LoadBalancer", "--name=my-service")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Service exposed successfully.")
}
