package main

import (
	"log"
	"net"

	studentsserver "students_db/pkg/students_server"
	"students_db/studentsdb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Create new gRPC server instance
	s := grpc.NewServer()
	srv := &studentsserver.StudentsServer{StudentsStore: make(map[string]*studentsdb.Student)}
	// Register gRPC server
	studentsdb.RegisterStudentsDatabaseServer(s, srv)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	// Listen on port 8080
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server is starting...")
	// Start gRPC server
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
