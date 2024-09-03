package planetmint

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/planetmint/planetmint-go/x/machine/types"
	"github.com/rddl-network/rddl-prometheus-exporter/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client *Client

func GetClient() *Client {
	if client == nil {
		cfg := config.GetConfig()
		client = newClient(cfg.PlanetmintHost)
	}
	return client
}

type Client struct {
	host string
}

func newClient(host string) *Client {
	return &Client{host: host}
}

func (pmc *Client) GetActiveDeviceCount() (count float64, err error) {
	conn, err := pmc.getGRPCConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	machineClient := types.NewQueryClient(conn)
	response, err := machineClient.ActiveTrustAnchorCount(
		context.Background(),
		&types.QueryActiveTrustAnchorCountRequest{},
	)
	if err != nil {
		return 0, fmt.Errorf("error while fetching active trust anchor count: %s", err.Error())
	}

	return float64(response.Count), nil
}

func (pmc *Client) GetActivatedDeviceCount() (count float64, err error) {
	conn, err := pmc.getGRPCConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	machineClient := types.NewQueryClient(conn)
	response, err := machineClient.ActivatedTrustAnchorCount(
		context.Background(),
		&types.QueryActivatedTrustAnchorCountRequest{},
	)
	if err != nil {
		return 0, fmt.Errorf("error while fetching activated trust anchor count: %s", err.Error())
	}

	return float64(response.Count), nil
}

func (pmc *Client) getGRPCConnection() (conn *grpc.ClientConn, err error) {
	return grpc.Dial(
		pmc.host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
}
