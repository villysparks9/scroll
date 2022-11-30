package bridgeabi_test

import (
	"math/big"
	"testing"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	bridge_abi "scroll-tech/bridge/abi"
)

func TestPackRelayMessageWithProof(t *testing.T) {
	assert := assert.New(t)

	l1MessengerABI, err := bridge_abi.L1MessengerMetaData.GetAbi()
	assert.NoError(err)

	proof := bridge_abi.IL1ScrollMessengerL2MessageProof{
		BlockHeight: big.NewInt(0),
		BatchIndex:  big.NewInt(0),
		MerkleProof: make([]byte, 0),
	}
	_, err = l1MessengerABI.Pack("relayMessageWithProof", common.Address{}, common.Address{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), make([]byte, 0), proof)
	assert.NoError(err)
}

func TestPackCommitBatch(t *testing.T) {
	assert := assert.New(t)

	l1RollupABI, err := bridge_abi.RollupMetaData.GetAbi()
	assert.NoError(err)

	txns := make([]bridge_abi.IZKRollupLayer2Transaction, 5)
	for i := 0; i < 5; i++ {
		txns[i] = bridge_abi.IZKRollupLayer2Transaction{
			Caller:   common.Address{},
			Target:   common.Address{},
			Nonce:    0,
			Gas:      0,
			GasPrice: big.NewInt(0),
			Value:    big.NewInt(0),
			Data:     make([]byte, 0),
			R:        big.NewInt(0),
			S:        big.NewInt(0),
			V:        0,
		}
	}

	header := bridge_abi.IZKRollupLayer2BlockHeader{
		BlockHash:   common.Hash{},
		ParentHash:  common.Hash{},
		BaseFee:     big.NewInt(0),
		StateRoot:   common.Hash{},
		BlockHeight: 0,
		GasUsed:     0,
		Timestamp:   0,
		ExtraData:   make([]byte, 0),
		Txs:         txns,
	}

	batch := bridge_abi.IZKRollupLayer2Batch{
		BatchIndex: 0,
		ParentHash: common.Hash{},
		Blocks:     []bridge_abi.IZKRollupLayer2BlockHeader{header},
	}

	_, err = l1RollupABI.Pack("commitBatch", batch)
	assert.NoError(err)
}

func TestPackFinalizeBatchWithProof(t *testing.T) {
	assert := assert.New(t)

	l1RollupABI, err := bridge_abi.RollupMetaData.GetAbi()
	assert.NoError(err)

	proof := make([]*big.Int, 10)
	instance := make([]*big.Int, 10)
	for i := 0; i < 10; i++ {
		proof[i] = big.NewInt(0)
		instance[i] = big.NewInt(0)
	}

	_, err = l1RollupABI.Pack("finalizeBatchWithProof", common.Hash{}, proof, instance)
	assert.NoError(err)
}

func TestPackRelayMessage(t *testing.T) {
	assert := assert.New(t)

	l2MessengerABI, err := bridge_abi.L2MessengerMetaData.GetAbi()
	assert.NoError(err)

	_, err = l2MessengerABI.Pack("relayMessage", common.Address{}, common.Address{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), make([]byte, 0))
	assert.NoError(err)
}