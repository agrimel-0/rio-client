package client

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/agrimel-0/rio-grpc"
)

func (c Client) GetGpioList() error {
	request := &pb.ClientRequest{}
	stream, err := c.client.GetGPIOList(context.Background(), request)
	if err != nil {
		return err
	}

	var ioPins []IoPin

	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
		}
		// log.Println(feature)

		receivedPin := IoPin{
			Alias:    feature.GPIOLineAlias,
			Line:     int(feature.GPIOLineOffset),
			GpioChip: feature.GPIOChip,
			Value:    int(feature.GPIOLineValue),
			AsOutput: feature.GPIOOutput,
		}

		ioPins = append(ioPins, receivedPin)
	}

	fmt.Println(ioPins)

	return nil

}
