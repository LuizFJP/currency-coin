package main

import (
	"context"
	"io"
	"testing"

	pb "LuizFJP/currency-coin-grpc-go/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestListCoins(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	client := pb.NewCurrencyCoinServiceClient(conn)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Success", func(mt *mtest.T) {
		collection = mt.Coll
		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "name", Value: "BitCoin"},
			{Key: "price", Value: 1.300},
			{Key: "vote", Value: 517},
		})
		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bson.D{
			{Key: "name", Value: "Ethereum"},
			{Key: "price", Value: 0.100},
			{Key: "vote", Value: 2500},
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)

		req := &pb.ListCoinRequest{}
		stream, err := client.ListCoins(context.Background(), req)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		count := 0

		for {
			_, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			count++
		}

		if count != 2 {
			t.Errorf("Expected 2, got: %d", count)
		}
	})
}
