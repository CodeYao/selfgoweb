package main

const (
	address     = "localhost:9092"
	defaultName = "abc"
)

// func main() {
// 	// Set up a connection to the server.
// 	conn, err := grpc.Dial(address, grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()
// 	c := pb.NewGreeterClient(conn)

// 	// Contact the server and print out its response.
// 	name := defaultName
// 	if len(os.Args) > 1 {
// 		name = os.Args[1]
// 	}
// 	r, err := c.AuthorityCtrl(context.Background(), &pb.AddressRequest{Addr: name})
// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}
// 	log.Printf("Greeting: %s", r.IsPermission)
// }
