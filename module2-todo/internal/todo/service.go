package todo

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/xxii22w/grpc-example/module2-todo/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	proto.UnimplementedTodoServiceServer
	tasks map[string]string
	lock  sync.Mutex
}

func NewService() *service {
	return &service{
		tasks: make(map[string]string),
		lock:  sync.Mutex{},
	}
}

func (s *service) AddTask(ctx context.Context, request *proto.AddTaskRequest) (*proto.AddTaskResponse, error) {
	if request.GetTask() == "" {
		return nil, status.Error(codes.InvalidArgument, "task connot be empty")
	}

	id := uuid.New().String()

	s.lock.Lock()
	s.tasks[id] = request.GetTask()
	s.lock.Unlock()

	return &proto.AddTaskResponse{
		Id: id,
	}, nil
}

func (s *service) CompleteTask(ctx context.Context, request *proto.CompleteTaskRequest) (*proto.CompleteTaskResponse, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	// check if task exist
	if _, ok := s.tasks[request.GetId()]; !ok {
		return nil, status.Error(codes.NotFound, "task not found")
	}

	// remove the task from store
	delete(s.tasks, request.GetId())

	// return the response
	return &proto.CompleteTaskResponse{}, nil
}

func (s *service) ListTasks(ctx context.Context, request *proto.ListTasksequest) (*proto.ListTasksResponse, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	// initialise a slice of tasks
	tasks := make([]*proto.Task, 0, len(s.tasks))

	// iterator through out map and populate slice
	for id, task := range s.tasks {
		tasks = append(tasks, &proto.Task{
			Id:   id,
			Task: task,
		})
	}

	// return the list task
	return &proto.ListTasksResponse{
		Tasks: tasks,
	}, nil
}
