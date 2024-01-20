package studentsserver

import (
	"context"
	"log"
	"strconv"

	"students_db/studentsdb"
	pb "students_db/studentsdb"
	"time"
)

type StudentsServer struct {
	StudentsStore map[string]*pb.Student
	studentsdb.UnimplementedStudentsDatabaseServer
}

// Handler for adding student
func (s *StudentsServer) AddStudent(ctx context.Context, req *pb.Student) (*pb.Student, error) {

	id := strconv.FormatInt(time.Now().Unix(), 12)
	req.Id = &id
	s.StudentsStore[id] = req
	log.Println("Student added: ", req.GetId(), req.GetName())
	return &pb.Student{Id: &id}, nil
}

// Handler for updating student
func (s *StudentsServer) UpdateStudent(ctx context.Context, req *pb.Student) (*pb.Student, error) {
	s.StudentsStore[req.GetId()] = req
	log.Println("Student updated: ", req.GetId(), req.GetName())
	return &pb.Student{Id: req.Id}, nil
}

// Handler for removing student
func (s *StudentsServer) DeleteStudent(ctx context.Context, req *pb.Student) (*pb.Student, error) {
	delete(s.StudentsStore, req.GetId())
	log.Println("Student deleted: ", req.GetId(), req.GetName())
	return &pb.Student{Id: req.Id}, nil
}

// Handler for getting student
func (s *StudentsServer) GetStudent(ctx context.Context, req *pb.Student) (*pb.Student, error) {
	for id, student := range s.StudentsStore {
		if student.GetId() == req.GetId() {
			return &pb.Student{Id: &id, Name: student.Name}, nil
		}
	}
	return &pb.Student{Id: req.Id}, nil
}
