package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/rdmtfl/order/data"
	"github.com/rdmtfl/order/proto/pb"
	"google.golang.org/grpc"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	Models data.Models
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Fatalf("failed to listen to gRPC: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterOrderServiceServer(s, &OrderServer{Models: app.Models})

	log.Printf("gRPC server listening on port %s", gRPCPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen to gRPC: %v", err)
	}
}

func (s *OrderServer) PostOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	order := req.GetOrderEntry()

	var args []string
	args = make([]string, 0)

	if strings.EqualFold(order.OrderId, "") {
		args = append(args, "Order ID is required")
	}
	if strings.EqualFold(order.CustomerId, "") {
		args = append(args, "Customer ID is required")
	}

	if len(args) > 0 {
		res := &pb.OrderResponse{
			Error:   true,
			Message: "missing fields",
			Fields:  args,
		}

		return res, nil
	}

	return nil, nil
	// orderEntry := data.OrderEntry{
	// 	OrderID:    input.OrderId,
	// 	CustomerID: input.CustomerId,
	// }

	// err := s.Models.OrderEntry.Insert(orderEntry)
	// if err != nil {
	// 	res := &pb.OrderResponse{
	// 		Result: "failed",
	// 	}
	// 	return res, err
	// }

	// // return repsonse
	// res := &pb.OrderResponse{
	// 	Result: "logged",
	// }
	// return res, nil

	// Convert the request to the internal data structure
	// orderEntry := data.OrderEntry{
	// 	OrderID:         req.OrderId,
	// 	CustomerID:      req.CustomerId,
	// 	OrderDate:       req.OrderDate.AsTime(),
	// 	Status:          req.Status,
	// 	Items:           convertItems(req.Items),
	// 	Subtotal:        req.Subtotal,
	// 	Tax:             req.Tax,
	// 	ShippingCost:    req.ShippingCost,
	// 	Total:           req.Total,
	// 	ShippingMethod:  req.ShippingMethod,
	// 	PaymentMethod:   req.PaymentMethod,
	// 	BillingAddress:  convertAddress(req.BillingAddress),
	// 	ShippingAddress: convertAddress(req.ShippingAddress),
	// }

	// err := s.Models.OrderEntry.Insert(orderEntry)
	// if err != nil {
	// 	return nil, err
	// }

	// return &pb.OrderResponse{
	// 	OrderId:      orderEntry.OrderID,
	// 	CreationDate: timestamppb.Now(),
	// }, nil
}
