package models

import (
	pb "github.com/Ilja-R/TeachMeSkillsHW/lesson-20/employee/proto"
)

func ModelToProto(e *Employee) *pb.Employee {
	if e == nil {
		return nil
	}
	return &pb.Employee{
		Id:    int32(e.ID),
		Name:  e.Name,
		Email: e.Email,
		Age:   int32(e.Age),
	}
}

func ProtoToModel(p *pb.Employee) *Employee {
	if p == nil {
		return nil
	}
	return &Employee{
		ID:    int(p.Id),
		Name:  p.Name,
		Email: p.Email,
		Age:   int(p.Age),
	}
}
