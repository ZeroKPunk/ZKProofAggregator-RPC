package aggregator

import (
	"context"
	aggregatorServer "github.com/ZKProofAggregator-RPC/aggregator/interface"
)

type Aggregator struct {
	aggregatorServer.UnimplementedAggregatorServiceServer
}

func (aggregator *Aggregator) GetSmtMerkleProof(ctx context.Context, input *aggregatorServer.SmtMerkleProofRequest) (*aggregatorServer.SmtMerkleProofResponse, error) {
	res := &aggregatorServer.SmtMerkleProofResponse{Proof: "", Hash: input.GetHash(), Result: 1}
	return res, nil
}
