package orchestrator

func (s *Server) GetUserDashboard(ctx context.Context, req *api.DashboardRequest) (*api.DashboardResponse, error) {
  userResp, _ := s.userClient.GetUser(ctx, &api.UserRequest{UserId: req.UserId})
  orderResp, _ := s.orderClient.GetOrders(ctx, &api.OrderRequest{UserId: req.UserId})

  return &api.DashboardResponse{
    Name:   userResp.Name,
    Orders: orderResp.Orders,
  }, nil
}

func main() {
  userConn, _ := grpc.Dial("user-service:50051", grpc.WithInsecure())
  orderConn, _ := grpc.Dial("order-service:50052", grpc.WithInsecure())

  server := &Server{
    userClient:  api.NewUserServiceClient(userConn),
    orderClient: api.NewOrderServiceClient(orderConn),
  }

  lis, _ := net.Listen("tcp", ":50050")
  grpcServer := grpc.NewServer()
  api.RegisterBFFServiceServer(grpcServer, server)
  grpcServer.Serve(lis)
}
