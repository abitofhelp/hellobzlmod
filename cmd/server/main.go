package main

//// server is used to implement helloworld.GreeterServer.
//type server struct {
//	v1.UnimplementedGreeterServer
//}

// SayHello implements helloworld.GreeterServer
//func (s *server) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
//
//	now := time.Now().UTC()
//	timestamp := timestamppb.Timestamp{
//		Seconds: now.Unix(),
//		Nanos:   int32(now.Nanosecond()),
//	}
//
//	log.Printf("Received: %v", in.GetName())
//	return &v1.HelloReply{Message: "Hello " + in.GetName(), RepliedUtc: &timestamp}, nil
//}

func main() {
	//port := 50051
	//lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//s := grpc.NewServer()
	//v1.RegisterGreeterServer(s, &server{})
	//log.Printf("server listening at %v", lis.Addr())
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
}
