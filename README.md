# Golang_GRPC_Implementation
Golang GRPC Implementation with sample codes for both client and server.

Following are the steps to download and install protobuff compiler.
It is developed by Google which helps to create .pb.go file

1) Download protobuff compiler from https://github.com/protocolbuffers/protobuf/releases
2) Create a proto extension file and write the following sample code

        syntax = "proto3";

        package proto;

        message Request{
          int64 a = 1;
          int64 b = 2;
        }

        message Response{
          int64 result = 1;
        }

        service AddService{
          rpc Add(Request) returns (Response);
          rpc Multiply(Request) returns (Response);
        }
  
3) Now to install grpc for golang write the following command "go get -u google.golang.org/grpc"

4) Now after download go package write the following command in the project root directory
protoc.exe --proto_path=proto --proto_path=google.golang.org/grpc --go_out=plugins=grpc:proto service.proto

5) After executing the command it creates a file named "service.pb.go". It has the schema and method and auto generated source code.

6) Now create a main.go file in the server folder and main.go in the client folder

7) Now in the server folder main.go write the following code

        package main
        
        // import the following packages

        import(
          "context"
          "proto"
          "net"
          "google.golang.org/grpc"
          "google.golang.org/grpc/reflection"
        )
        
        // create a struct with name server

        type server struct{}
        
        // main function

        func main(){
        
            // listening to tcp port 4040

            listener, err := net.Listen("tcp", ":4040")

            if err != nil{
              panic(err)
            }
            
            // hosting grpc server, it is insecure server

            srv := grpc.NewServer()
            
            // Registering Add Service Method

            proto.RegisterAddServiceServer(srv, &server{})

            reflection.Register(srv)

            if e:= srv.Serve(listener); e != nil{
              panic(err)
            }

        }
        
        // creating Add method

        func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error){
        
            // getting the value of A and B

            a, b := request.GetA(), request.GetB()
            
            // adding them and return the value to the response object of the GRPC

            result := a + b

            return &proto.Response{Result: result}, nil

        }
        
        // declaring multiply method

        func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error){
        
            // getting the value of A and B

            a, b := request.GetA(), request.GetB()
            
            // multiplying a and b

            result := a * b
            
            // returning the result to the GRPC response object

            return &proto.Response{Result: result}, nil

        }
        
 ###############################################################################################################################
 
 8) Create the client folder main.go and write the following code
 
        package main
        
        // importing the packages

        import(
          "context"
          "log"
          "time"
          "google.golang.org/grpc"
          pb "proto"
        )
        
        // creating a constant with tcp address 

        const (
          address = "localhost:4040"
        )
        
        // declaring main method

        func main(){
        
          // creating a connect object and connecting to tcp server

          // Set up a connection to the server.
          conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
          
          if err != nil {
            log.Fatalf("did not connect: %v", err)
          }
          defer conn.Close()
          
          // invoking the client service to connect to grpc server

          c := pb.NewAddServiceClient(conn)

          // Contact the server and print out its response.
          
          // setting timeout for go channels

          ctx, cancel := context.WithTimeout(context.Background(), time.Second)

          defer cancel()
          
          // adding values to Add method

          r, err := c.Add(ctx, &pb.Request{A: 5, B:2})

          if err != nil {
            log.Fatalf("could not greet: %v", err)
          }
          
          // getting result from server

          log.Printf("Greeting: %s", r.GetResult())
          
          // adding values to Multiply method

          r, err = c.Multiply(ctx, &pb.Request{A: 5, B:6})

          if err != nil {
            log.Fatalf("could not greet: %v", err)
          }
          
          // getting result from server

          log.Printf("Greeting: %s", r.GetResult())
        }
        
There are more method to create ssl secure connections. Follow the original documentation in the following link. 
https://github.com/grpc/grpc-go
