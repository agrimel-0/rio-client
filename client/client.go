package client

import (
	pb "github.com/agrimel-0/rio-grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client pb.RioClient // Rio client instance
}

func Start(serverAddr *string) (Client, error) {

	logrus.Info("starting client...")

	client := Client{}

	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Error("error dialing: %v", err)
		return client, err
	}
	// defer conn.Close()

	client.client = pb.NewRioClient(conn)

	logrus.Debug("rio client created")

	return client, nil

}
