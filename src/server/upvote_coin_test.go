package main

import (
	"context"
	"testing"

	pb "LuizFJP/currency-coin-grpc-go/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

)

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewCurrencyCoinServiceClient(conn)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Success to upvote", func(mt *mtest.T) {
		collection = mt.Coll
		fakeCoin := &CoinItem{
			Name: "BTC",
			Price: 1.300000,
			Vote: 2500,
		}

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
		})

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "name", Value: "BTC"},
			{Key: "price", Value: 1.300000},
			{Key: "vote", Value: 2501},
		}))

		req := &pb.CoinRequest{Name: fakeCoin.Name}
		_, err := c.UpvoteCoin(context.Background(), req)

		if err != nil {
			t.Errorf("Something went wrong: %v", err)
		}
	})
}

