package main

// import "google.golang.org/grpc"
import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	pb "grpcdemo/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "Jansen"
)

var (
	name     = flag.String("name", defaultName, "Name to greet")
	todo     = flag.String("item", "do nothing", "Item to add")
	getTodos = flag.Bool("get", false, "Get todo items")
)

func main() {
	flag.Parse()
	addr := "localhost:50051"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// c := pb.NewGreetClient(conn)
	c := pb.NewTodoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if *getTodos {
		stream, err := c.GetTodos(ctx, &pb.Empty{})
		if err != nil {
			log.Fatalf("could not make: %v", err)
		}

		for {
			item, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("problems receiving from stream. %v", err)
			}
			log.Printf("Item: %v", item)
		}
		return
	}
	r, err := c.MakeTodo(ctx, &pb.TodoItem{
		Id:   0,
		Item: *todo,
	})

	if err != nil {
		log.Fatalf("could not make: %v", err)
	}

	log.Printf("Made Todo item: %v", r)
}
