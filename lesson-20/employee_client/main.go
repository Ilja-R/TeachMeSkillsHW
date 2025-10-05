package main

import (
	"context"
	pb "github.com/Ilja-R/TeachMeSkillsHW/lesson-20/employee/proto"
	converters "github.com/Ilja-R/TeachMeSkillsHW/lesson-20/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewEmployeeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// CreateEmployee
	req := &pb.CreateEmployeeRequest{
		Name:  "John Doe",
		Email: "john_doe@test.com",
		Age:   30,
	}
	res, err := client.CreateEmployee(ctx, req)
	if err != nil {
		log.Fatalf("could not create employee: %v", err)
	}
	log.Printf("Created Employee: %+v", converters.ProtoToModel(res.Employee))

	// CreateEmployee
	req = &pb.CreateEmployeeRequest{
		Name:  "Alice Smith",
		Email: "alice@example.com",
		Age:   28,
	}
	res, err = client.CreateEmployee(ctx, req)
	if err != nil {
		log.Fatalf("could not create employee: %v", err)
	}
	log.Printf("Created Employee: %+v", converters.ProtoToModel(res.Employee))

	// UpdateEmployee
	updateReq := &pb.UpdateEmployeeRequest{
		Id:    res.Employee.Id,
		Name:  "Alice Johnson",
		Email: "alice@example.com",
		Age:   29,
	}
	updateRes, err := client.UpdateEmployeeByID(ctx, updateReq)
	if err != nil {
		log.Fatalf("could not update employee: %v", err)
	}
	log.Printf("Updated Employee: %+v", converters.ProtoToModel(updateRes.Employee))

	// GetEmployeeByID
	getReq := &pb.GetEmployeeRequest{Id: res.Employee.Id}
	getRes, err := client.GetEmployeeByID(ctx, getReq)
	if err != nil {
		log.Fatalf("could not get employee: %v", err)
	}
	log.Printf("Fetched Employee by ID: %+v", converters.ProtoToModel(getRes.Employee))

	// GetAllEmployees
	allRes, err := client.GetAllEmployees(ctx, &pb.GetAllEmployeesRequest{})
	if err != nil {
		log.Fatalf("could not get all employees: %v", err)
	}
	log.Println("All Employees:")
	for _, emp := range allRes.Employees {
		log.Printf("Employee: %+v", converters.ProtoToModel(emp))
	}
	log.Printf("Total Employees: %d", len(allRes.Employees))

	// DeleteEmployee
	delReq := &pb.DeleteEmployeeRequest{Id: res.Employee.Id}
	_, err = client.DeleteEmployeeByID(ctx, delReq)
	if err != nil {
		log.Fatalf("could not delete employee: %v", err)
	}
	log.Printf("Deleted Employee with ID: %d", res.Employee.Id)
}
