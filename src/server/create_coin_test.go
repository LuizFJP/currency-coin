package main

import (
	"context"
	"testing"

	pb "LuizFJP/currency-coin-grpc-go/proto"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

)

func TestCreateCoin(t *testing.T) {
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
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		fakeCoin := &pb.CreateCoinRequest{
			Name: "Ethereum",
			Price:    0.0006,
		}

	req := &pb.CreateCoinRequest{Name: fakeCoin.Name, Price: fakeCoin.Price}
	_, err := client.CreateCoin(context.Background(), req)

	if err != nil {
		t.Errorf("HelloTest(%v) got unexpected error", err)
}

},
	)}
