package main

import (
	"fmt"
	"net"
	"time"
)

type Container struct {
	Host string
	Port string
}

type TCPHealthChecker struct {
	Timeout time.Duration
}

func (t *TCPHealthChecker) IsHealthy(container Container) (bool, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(container.Host, container.Port), t.Timeout)
	if err != nil {
		return false, fmt.Errorf("TCP dial error: %w", err)
	}
	defer conn.Close()

	return true, nil
}

func main() {
	checker := &TCPHealthChecker{
		Timeout: 1 * time.Second,
	}

	container := Container{
		Host: "localhost",
		Port: "8080",
	}

	healthy, err := checker.IsHealthy(container)
	if err != nil {
		fmt.Println("Health check error:", err)
		return
	}

	if healthy {
		fmt.Println("Service is healthy")
	} else {
		fmt.Println("Service is not healthy")
	}
}
