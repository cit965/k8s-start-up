package main

import "fmt"

type ReplicationManager struct {
	// ...
}

type Status string

const (
	PodRunning Status = "Running"
	PodPending Status = "Pending"
	PodStop    Status = "Stop"
)

type Pod struct {
	Status Status
}

func (rm *ReplicationManager) filterActivePods(podList []Pod) []Pod {
	var activePods []Pod
	for _, pod := range podList {
		if pod.Status == PodRunning {
			activePods = append(activePods, pod)
		}
	}
	return activePods
}

func main() {
	rm := &ReplicationManager{}

	// 创建一些测试数据
	podList := []Pod{
		{Status: PodRunning},
		{Status: PodPending},
		{Status: PodStop},
		{Status: PodRunning},
		{Status: PodPending},
		{Status: PodRunning},
	}

	// 使用filterActivePods方法筛选出活跃的Pod
	activePods := rm.filterActivePods(podList)

	// 输出筛选结果
	fmt.Printf("Active Pods: %v\n", len(activePods))
	for i, pod := range activePods {
		fmt.Printf("Pod %d: %v\n", i+1, pod.Status)
	}
}
