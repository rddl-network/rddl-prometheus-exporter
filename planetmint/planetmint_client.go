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

var client *PlanetmintClient

func GetClient() *PlanetmintClient {
	if client == nil {
		cfg := config.GetConfig()
		client = newPlanetmintClient(cfg.PlanetmintHost)
	}
	return client
}

type PlanetmintClient struct {
	host string
}

func newPlanetmintClient(host string) *PlanetmintClient {
	return &PlanetmintClient{host: host}
}

func (pmc *PlanetmintClient) GetActiveDeviceCount() (count float64, err error) {
	conn, err := pmc.getGRPCConnection()
	if err != nil {
		return
	}

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

func (pmc *PlanetmintClient) GetActivatedDeviceCount() (count float64, err error) {
	conn, err := pmc.getGRPCConnection()
	if err != nil {
		return
	}

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

func (pmc *PlanetmintClient) getGRPCConnection() (conn *grpc.ClientConn, err error) {
	return grpc.Dial(
		pmc.host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
}
