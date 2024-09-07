package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/nshafer/phx"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	log.Println("Starting the application...")
	clusterId := os.Getenv("CLUSTER_ID")
	if clusterId == "" {
		log.Fatal("CLUSTER_ID environment variable is not set")
	}

	secret := os.Getenv("CLUSTER_SECRET")
	if secret == "" {
		log.Fatal("CLUSTER_SECRET environment variable is not set")
	}

	endpointURL := os.Getenv("ENDPOINT_URL")
	if endpointURL == "" {
		endpointURL = "wss://ranching.farm/socket/kubernetes/cluster"
	}

	log.Printf("Connecting to WebSocket at %s", endpointURL)
	endPoint, err := url.Parse(endpointURL)
	if err != nil {
		log.Fatal("Failed to parse WebSocket URL:", err)
	}

	socket := phx.NewSocket(endPoint)
	err = socket.Connect()
	if err != nil {
		log.Fatal("Failed to connect to socket:", err)
	}
	log.Println("Successfully connected to WebSocket")

	log.Printf("Joining channel: cluster:%s", clusterId)
	channel := socket.Channel(fmt.Sprintf("cluster:%s", clusterId), nil)
	join, err := channel.Join()
	if err != nil {
		log.Fatal("Failed to join channel:", err)
	}

	join.Receive("ok", func(response any) {
		log.Println("Joined channel:", channel.Topic(), response)
		sendClusterInfo(channel)
	})

	channel.On("cmd", func(payload any) {
		handleCommand(channel, payload)
	})

	log.Println("Main loop started. Waiting for events...")
	select {} // Keep the program running
}

func sendClusterInfo(channel *phx.Channel) {
	log.Println("Sending cluster info...")
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Println("Failed to get in-cluster config:", err)
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println("Failed to create Kubernetes client:", err)
		return
	}

	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Println("Failed to list nodes:", err)
		return
	}

	log.Printf("Found %d nodes in the cluster", len(nodes.Items))
	nodeInfo := make([]map[string]string, 0)
	for _, node := range nodes.Items {
		nodeInfo = append(nodeInfo, map[string]string{
			"name":   node.Name,
			"status": string(node.Status.Phase),
		})
		log.Printf("Node: %s, Status: %s", node.Name, node.Status.Phase)
	}

	clusterInfo := map[string]interface{}{
		"nodes": nodeInfo,
	}

	push, err := channel.Push("info", clusterInfo)
	if err != nil {
		log.Println("Failed to send cluster info:", err)
		return
	}

	push.Receive("ok", func(response any) {
		log.Println("Cluster info sent successfully:", response)
	})
}

func handleCommand(channel *phx.Channel, payload any) {
	log.Println("Received command payload:", payload)
	cmd, ok := payload.(map[string]interface{})
	if !ok {
		log.Println("Invalid payload format")
		return
	}

	command, ok := cmd["command"].(string)
	if !ok {
		log.Println("Invalid command format")
		return
	}

	log.Printf("Received command: %+v\n", cmd)

	args, ok := cmd["arguments"].(string)
	if !ok {
		log.Println("Invalid args format")
		return
	}

	uuid, ok := cmd["uuid"].(string)
	if !ok {
		log.Println("Invalid uuid format")
		return
	}

	output, err := executeCommand(command, args)
	if err != nil {
		log.Printf("Error executing command: %v", err)
		output = fmt.Sprintf("Error: %v", err)
	} else {
		log.Println("Command executed successfully")
	}

	response := map[string]interface{}{
		"command":   command,
		"arguments": args,
		"output":    output,
		"uuid":      uuid,
	}

	push, err := channel.Push("output", response)
	if err != nil {
		log.Println("Failed to send command output:", err)
		return
	}

	push.Receive("ok", func(response any) {
		log.Println("Command output sent successfully:", response)
	})
}

func executeCommand(command string, params string) (string, error) {
	log.Printf("Executing command: %s %s", command, params)

	// Split the command string into arguments
	args := strings.Fields(params)
	cmd := exec.Command(command, args...)

	// Capture both stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Execute the command
	err := cmd.Run()

	if err != nil {
		return fmt.Sprintf("Error: %v\nStdout: %s\nStderr: %s", err, stdout.String(), stderr.String()), err
	}

	log.Println("Command executed successfully")
	return stdout.String(), nil
}
