package main

import (
	"context"
	"flag"
	"fmt"
	pb "grpcdemo/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGreetServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Greeting: "Hello " + in.GetName()}, nil
}

type todoserver struct {
	pb.UnimplementedTodoServer
	todos []*pb.TodoItem
}

func (s *todoserver) MakeTodo(ctx context.Context, in *pb.TodoItem) (*pb.TodoItem, error) {
	log.Printf("Received: %v", in)
	item := pb.TodoItem{
		Id:   int32(len(s.todos)),
		Item: in.GetItem(),
	}

	s.todos = append(s.todos, &item)

	return &item, nil
}

func (s *todoserver) GetTodos(_ *pb.Empty, stream pb.Todo_GetTodosServer) error {
	for _, todo := range s.todos {
		stream.Send(todo)
	}
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// pb.RegisterGreetServer(s, &server{})
	pb.RegisterTodoServer(s, &todoserver{todos: []*pb.TodoItem{}})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
