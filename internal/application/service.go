package application

import (
	"github.com/ark-network/ark/common/bitcointree"
	"github.com/ark-network/ark/pkg/client-sdk/client"
	grpcclient "github.com/ark-network/ark/pkg/client-sdk/client/grpc"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

type Service struct {
	privateKey *secp256k1.PrivateKey

	arkClient client.TransportClient
	sessions  map[string]bitcointree.SignerSession
}

func New(serverURL string) (*Service, error) {
	arkClient, err := grpcclient.NewClient(serverURL)
	if err != nil {
		return nil, err
	}
	return &Service{
		arkClient: arkClient,
	}, nil
}

func (s *Service) Start() error {
	// start listening for GetEventStream
	// handle events
	return nil
}

func (s *Service) Stop() error {
	return nil
}

func (s *Service) handleRoundSigningStartedEvent(event *client.RoundSigningStartedEvent) {
	// check if validator's key is in the round
	// intantiace a new TreeSignerSession -> set it in the sessions map
	// = bitcointree.NewTreeSignerSession(s.privateKey....)
	// generate nonces
	// submit nonces
}

func (s *Service) handleRoundSigningNoncesGeneratedEvent(event *client.RoundSigningNoncesGeneratedEvent) {
	// get the session from the sessions map
	// sign the tree
	// submit tree signatures
}
