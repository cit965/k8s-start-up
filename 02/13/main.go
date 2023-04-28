package main

import (
	"fmt"
	"testing"
)

type Deployment struct {
	DesiredState struct {
		Replicas int
	}
}

type clientInterface interface {
	GetDeployment(string) (*Deployment, error)
	UpdateDeployment(*Deployment) (*Deployment, error)
}

func ScaleDeployment(name string, replicas int, client clientInterface) error {
	deployment, err := client.GetDeployment(name)
	if err != nil {
		return err
	}
	deployment.DesiredState.Replicas = replicas
	deploymentOut, err := client.UpdateDeployment(deployment)
	if err != nil {
		return err
	}
	fmt.Print(deploymentOut)
	return nil
}

type FakeKubeClient struct {
	actions []struct {
		action string
		value  interface{}
	}
}

func (f *FakeKubeClient) GetDeployment(name string) (*Deployment, error) {
	f.actions = append(f.actions, struct {
		action string
		value  interface{}
	}{"get-deployment", name})
	return &Deployment{}, nil
}

func (f *FakeKubeClient) UpdateDeployment(d *Deployment) (*Deployment, error) {
	f.actions = append(f.actions, struct {
		action string
		value  interface{}
	}{"update-deployment", d})
	return d, nil
}

func TestScaleDeployment(t *testing.T) {
	fakeClient := FakeKubeClient{}
	name := "name"
	replicas := 17
	ScaleDeployment(name, replicas, &fakeClient)
	if len(fakeClient.actions) != 2 {
		t.Errorf("Unexpected actions: %#v", fakeClient.actions)
	}
	if fakeClient.actions[0].action != "get-deployment" ||
		fakeClient.actions[0].value.(string) != name {
		t.Errorf("Unexpected action: %#v", fakeClient.actions[0])
	}
	deployment := fakeClient.actions[1].value.(*Deployment)
	if fakeClient.actions[1].action != "update-deployment" ||
		deployment.DesiredState.Replicas != 17 {
		t.Errorf("Unexpected action: %#v", fakeClient.actions[1])
	}
}

// 在这个示例中，ScaleDeployment函数从客户端接口获取部署对象，更新其副本数量，然后使用客户端接口将更新后的部署对象发送回服务器。TestScaleDeployment函数测试了客户端接口是否调用了正确的操作，并验证了更新后的部署副本数量。
