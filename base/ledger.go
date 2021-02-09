package base

import (
	"encoding/hex"
	"fmt"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

func (c *Client) LedgerQueryConfig(
	options ...ledger.RequestOption,
) (fab.ChannelCfg, error) {
	return c.ledgerClient.QueryConfig(options...)
}

func (c *Client) LedgerQueryConfigBlock(
	options ...ledger.RequestOption,
) (*common.Block, error) {
	return c.ledgerClient.QueryConfigBlock(options...)
}

func (c *Client) LedgerQueryInfo(
	options ...ledger.RequestOption,
) (*fab.BlockchainInfoResponse, error) {
	return c.ledgerClient.QueryInfo(options...)
}

func (c *Client) LedgerQueryBlock(
	blockNumber uint64,
	options ...ledger.RequestOption,
) (*common.Block, error) {
	return c.ledgerClient.QueryBlock(blockNumber, options...)
}

func (c *Client) LedgerQueryBlockByHash(
	hash string,
	options ...ledger.RequestOption,
) (*common.Block, error) {
	blockHash, err := hex.DecodeString(hash)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hash: %v", err)
	}
	return c.ledgerClient.QueryBlockByHash(blockHash, options...)
}

func (c *Client) LedgerQueryBlockByTxID(
	txID string,
	options ...ledger.RequestOption,
) (*common.Block, error) {
	transactionID := fab.TransactionID(txID)
	return c.ledgerClient.QueryBlockByTxID(transactionID, options...)
}

func (c *Client) LedgerQueryTransaction(
	txID string,
	options ...ledger.RequestOption,
) (*peer.ProcessedTransaction, error) {
	transactionID := fab.TransactionID(txID)
	return c.ledgerClient.QueryTransaction(transactionID, options...)
}

func (c *Client) LedgerQueryBlockchainInfo(
	options ...ledger.RequestOption,
) (*BlockchainInfo, error) {
	respFrom, err := c.LedgerQueryInfo(options...)
	if err != nil {
		return nil, fmt.Errorf("failed to call LedgerQueryInfo: %v", err)
	}
	blockchainInfo := respFrom.BCI
	if blockchainInfo == nil {
		return nil, fmt.Errorf("nil blockchain info")
	}
	respTo, err := DecodeBlockchainInfo(blockchainInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to decode blockchainInfo(%+v): %v", blockchainInfo, err)
	}
	return respTo, nil
}

func (c *Client) LedgerQueryDecodedBlock(
	blockNumber uint64,
	options ...ledger.RequestOption,
) (*Block, error) {
	respFrom, err := c.LedgerQueryBlock(blockNumber, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to call LedgerQueryBlock: %v", err)
	}
	respTo, err := DecodeBlock(respFrom)
	if err != nil {
		return nil, fmt.Errorf("failed to decode block: %v", err)
	}
	return respTo, nil
}

func (c *Client) LedgerQueryDecodedBlockByHash(
	hash string,
	options ...ledger.RequestOption,
) (*Block, error) {
	respFrom, err := c.LedgerQueryBlockByHash(hash, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to call LedgerQueryBlockByHash: %v", err)
	}
	respTo, err := DecodeBlock(respFrom)
	if err != nil {
		return nil, fmt.Errorf("failed to decode block: %v", err)
	}
	return respTo, nil
}

func (c *Client) LedgerQueryDecodedBlockByTxID(
	txID string,
	options ...ledger.RequestOption,
) (*Block, error) {
	respFrom, err := c.LedgerQueryBlockByTxID(txID, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to call LedgerQueryBlockByTxID: %v", err)
	}
	respTo, err := DecodeBlock(respFrom)
	if err != nil {
		return nil, fmt.Errorf("failed to decode block: %v", err)
	}
	return respTo, nil
}

func (c *Client) LedgerQueryDecodedTransaction(
	txID string,
	options ...ledger.RequestOption,
) (*ProcessedTransaction, error) {
	respFrom, err := c.LedgerQueryTransaction(txID, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to call LedgerQueryTransaction: %v", err)
	}
	respTo, err := DecodeProcessedTransaction(respFrom)
	if err != nil {
		return nil, fmt.Errorf("failed to decode processedTransaction: %v", err)
	}
	return respTo, nil
}
