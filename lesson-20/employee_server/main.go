package main

import (
	"context"
	"fmt"
	pb "github.com/Ilja-R/TeachMeSkillsHW/lesson-20/employee/proto"
	"github.com/Ilja-R/TeachMeSkillsHW/lesson-20/models"
	converters "github.com/Ilja-R/TeachMeSkillsHW/lesson-20/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"sync"
)

type server struct {
	pb.UnimplementedEmployeeServiceServer

	// in-memory storage for employees
	mu        sync.RWMutex
	currentID int
	employees map[int32]*models.Employee
}

func (s *server) GetAllEmployees(_ context.Context, _ *pb.GetAllEmployeesRequest) (*pb.GetAllEmployeesResponse, error) {
	s.mu.RLock()
	employees := make([]*pb.Employee, 0, len(s.employees))
	for _, emp := range s.employees {
		employees = append(employees, converters.ModelToProto(emp))
	}
	s.mu.RUnlock()

	return &pb.GetAllEmployeesResponse{
		Employees: employees,
	}, nil
}

func (s *server) GetEmployeeByID(_ context.Context, req *pb.GetEmployeeRequest) (*pb.EmployeeResponse, error) {
	s.mu.RLock()
	emp, exists := s.employees[req.Id]
	s.mu.RUnlock()
	if !exists {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("employee with ID %d not found", req.Id))
	}
	return &pb.EmployeeResponse{
		Employee: converters.ModelToProto(emp),
	}, nil
}

func (s *server) CreateEmployee(_ context.Context, req *pb.CreateEmployeeRequest) (*pb.EmployeeResponse, error) {
	s.mu.Lock()
	s.currentID++
	newEmp := &models.Employee{
		ID:    s.currentID,
		Name:  req.Name,
		Email: req.Email,
		Age:   int(req.Age),
	}
	s.employees[int32(newEmp.ID)] = newEmp
	s.mu.Unlock()

	return &pb.EmployeeResponse{
		Employee: converters.ModelToProto(newEmp),
	}, nil
}

func (s *server) UpdateEmployeeByID(_ context.Context, req *pb.UpdateEmployeeRequest) (*pb.EmployeeResponse, error) {
	s.mu.Lock()
	emp, exists := s.employees[req.Id]
	if !exists {
		s.mu.Unlock()
		return nil, status.Error(codes.NotFound, fmt.Sprintf("employee with ID %d not found", req.Id))
	}
	emp.Name = req.Name
	emp.Email = req.Email
	emp.Age = int(req.Age)
	s.employees[req.Id] = emp
	s.mu.Unlock()

	return &pb.EmployeeResponse{
		Employee: converters.ModelToProto(emp),
	}, nil
}

func (s *server) DeleteEmployeeByID(_ context.Context, req *pb.DeleteEmployeeRequest) (*pb.EmployeeResponse, error) {
	s.mu.Lock()
	emp, exists := s.employees[req.Id]
	if !exists {
		s.mu.Unlock()
		return nil, status.Error(codes.NotFound, fmt.Sprintf("employee with ID %d not found", req.Id))
	}
	delete(s.employees, req.Id)
	s.mu.Unlock()
	return &pb.EmployeeResponse{
		Employee: converters.ModelToProto(emp),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(s, &server{employees: make(map[int32]*models.Employee)})

	log.Printf("gRPC empoyee server started!")
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
