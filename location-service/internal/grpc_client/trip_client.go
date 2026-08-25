package grpc_client

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/farmcreepissohard/Ride-hailing-sytem/pkg/pb"
	"google.golang.org/grpc"
)

type TripGrpcClient interface {
	MatchTrip(ctx context.Context, tripId string, driverId string) (bool, error)
	StartTrip(ctx context.Context, tripId string, driverId string) (bool, error)
	CompleteTrip(ctx context.Context, tripId string, driverId string) (bool, error)
	CancelTrip(ctx context.Context, tripId string, cancelledBy string, reason string) (bool, error)
}

type tripGrpcClient struct {
	client pb.UpdateTripInformationClient
	conn   *grpc.ClientConn
}

func NewTripGrpcClient(client pb.UpdateTripInformationClient,
	conn *grpc.ClientConn) TripGrpcClient {
	return &tripGrpcClient{client: client, conn: conn}
}

func (tripGrpc *tripGrpcClient) handler(res *pb.TripResponse, err error) (bool, error) {
	if err != nil {
		log.Printf("Failed to call MatchTrip gRPC: %v", err)
		return false, err
	}

	if res.Status == pb.ResponseStatus_SUCCESS {
		return true, nil
	}

	return false, errors.New(res.GetMessage())
}

func (tripGrpc *tripGrpcClient) MatchTrip(ctx context.Context, tripId string, driverId string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req := pb.TripActionRequest{
		TripId:   tripId,
		DriverId: driverId,
	}

	res, err := tripGrpc.client.MatchTrip(ctx, &req)
	return tripGrpc.handler(res, err)
}

func (tripGrpc *tripGrpcClient) StartTrip(ctx context.Context, tripId string, driverId string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req := pb.TripActionRequest{
		TripId:   tripId,
		DriverId: driverId,
	}

	res, err := tripGrpc.client.StartTrip(ctx, &req)
	return tripGrpc.handler(res, err)
}

func (tripGrpc *tripGrpcClient) CompleteTrip(ctx context.Context, tripId string, driverId string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req := pb.TripActionRequest{
		TripId:   tripId,
		DriverId: driverId,
	}

	res, err := tripGrpc.client.CompleteTrip(ctx, &req)
	return tripGrpc.handler(res, err)
}

func (tripGrpc *tripGrpcClient) CancelTrip(ctx context.Context, tripId string, cancelledBy string, reason string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req := pb.CancelTripRequest{
		TripId:        tripId,
		IdCancelledBy: cancelledBy,
		Reason:        reason,
	}

	res, err := tripGrpc.client.CancelTrip(ctx, &req)
	return tripGrpc.handler(res, err)
}
