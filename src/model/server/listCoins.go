package main

import (
	pb "github.com/LuizFJP/currency-coin-grpc-go"
)

func (s *Server) ListCoins(ctx context.Context, *empty.Empty) (*CoinResponse, error)