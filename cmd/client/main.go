package main

func main() {

	//now := time.Now().UTC()
	//timestamp := timestamppb.Timestamp{
	//	Seconds: now.Unix(),
	//	Nanos:   int32(now.Nanosecond()),
	//}
	//
	//req := v1.HelloRequest{
	//	Name:    "Mikie",
	//	SentUtc: &timestamp,
	//}
	//
	//conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	//c := v1.NewGreeterClient(conn)
	//
	//// Contact the server and print out its response.
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//r, err := c.SayHello(ctx, &req)
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.GetMessage())

}
