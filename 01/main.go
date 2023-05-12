package main

import (
	"github.com/gin-gonic/gin"
)

// 创建任务的 API 处理函数
func createTaskHandler(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}
	taskRegistry := TaskRegistry{}
	err := taskRegistry.CreateTask(task)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(201, gin.H{"message": "Task created successfully"})
}

// 获取任务列表的 API 处理函数
func listTasksHandler(c *gin.Context) {

	taskRegistry := TaskRegistry{}
	tasks, err := taskRegistry.ListTasks()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to list tasks"})
		return
	}

	c.JSON(200, tasks)
}

func listServicesHandler(c *gin.Context) {
	c.JSON(200, nil)
}

func createServicesHandler(c *gin.Context) {
	c.JSON(200, nil)
}

func main() {
	// 初始化 Gin Web 框架
	r := gin.Default()

	// 注册 API 路由
	//r.POST("/tasks", createTaskHandler)
	//r.GET("/tasks", listTasksHandler)
	//r.GET("/services", listServicesHandler)
	//r.POST("/services", createServicesHandler)
	r.Any("resource/:type", restHandler)

	// 启动 Web 服务器
	r.Run(":8080")
}

// Task 结构体
type Task struct {
	Name        string
	Description string
}

type Registry interface {
	ListResource() ([]interface{}, error)
	CreateResource(task interface{}) error
}

type MysqlTaskRegistry struct {
}

func (t *MysqlTaskRegistry) ListTasks() ([]Task, error) {
	return []Task{{
		Name:        "mysql",
		Description: "test",
	}}, nil
}

func (t *MysqlTaskRegistry) CreateTask(task Task) error {
	return nil
}

type MockTaskRegistry struct {
}

func (t *MockTaskRegistry) ListTasks() ([]interface{}, error) {
	return []Task{{
		Name:        "test",
		Description: "test",
	}}, nil
}

func (t *MockTaskRegistry) CreateTask(task interface{}) error {
	return nil
}

type handerStorage interface {
	List(c *gin.Context)
	Create(c *gin.Context)
}

type TaskStorage struct {
}

func (t *TaskStorage) List(c *gin.Context) {
	taskRegistry := MockTaskRegistry{}
	tasks, err := taskRegistry.ListTasks()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to list tasks"})
		return
	}

	c.JSON(200, tasks)
}

func (t *TaskStorage) Create(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}
	taskRegistry := MockTaskRegistry{}
	err := taskRegistry.CreateTask(task)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(201, gin.H{"message": "Task created successfully"})
}

type ServceStorage struct {
}

func (t *ServceStorage) List(c *gin.Context) {

	c.JSON(200, nil)
}

func (t *ServceStorage) Create(c *gin.Context) {
	c.JSON(201, gin.H{"message": "service created successfully"})
}

func restHandler(c *gin.Context) {

	m := map[string]handerStorage{"task": &TaskStorage{}, "service": &ServceStorage{}}

	resourceType := c.Param("type")
	switch c.Request.Method {
	case "GET":
		m[resourceType].List(c)
	case "POST":
		m[resourceType].Create(c)
	}
}
