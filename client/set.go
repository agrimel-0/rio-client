package client

import (
	"context"

	pb "github.com/agrimel-0/rio-grpc"
	"github.com/sirupsen/logrus"
)

func (c Client) SetByOffset(lineOffset, state int) error {

	response, err := c.client.SetGPIObyOffset(context.Background(), &pb.GPIOselected{GPIOLineOffset: int32(lineOffset), GPIOLineValue: int32(state)})

	if err != nil {
		return err
	}

	logrus.Info(response)

	return nil
}
