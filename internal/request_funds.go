package internal

import (
	"context"
	"errors"

	faucetpb "github.com/Liangyu-Zhou/registry-demo/proto/faucet"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const weiPerETH = 1e18

// RequestFunds from an Ethereum faucet. Requires a valid captcha response.
func (s *Server) RequestFunds(
	ctx context.Context, req *faucetpb.FundingRequest,
) (*faucetpb.FundingResponse, error) {
	if req.WalletAddress == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Request needs a valid ETH wallet address")
	}
	ipAddress, err := s.getIPAddress(ctx)
	if err != nil {
		log.WithError(err).Error("Could not fetch IP from request")
		return nil, status.Errorf(codes.FailedPrecondition, "Could not get IP address from request: %v", err)
	}

	log.WithFields(logrus.Fields{
		"ipAddress": ipAddress,
		"address":   req.WalletAddress,
	}).Info("Attempting to fund address")

	return &faucetpb.FundingResponse{
		Amount:          "100",
		TransactionHash: "txHash",
	}, nil
}

func (s *Server) getIPAddress(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("x-forwarded-for")) < 1 {
		return "", errors.New("metadata not ok")
	}
	address := md.Get("x-forwarded-for")[0]
	return address, nil
}
