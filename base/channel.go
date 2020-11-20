package base

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (c *Client) ChannelQuery(
	request channel.Request,
	options ...channel.RequestOption,
) (channel.Response, error) {
	return c.channelClient.Query(request, options...)
}

func (c *Client) ChannelExecute(
	request channel.Request,
	options ...channel.RequestOption,
) (channel.Response, error) {
	return c.channelClient.Execute(request, options...)
}
