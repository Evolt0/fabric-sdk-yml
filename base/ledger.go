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
