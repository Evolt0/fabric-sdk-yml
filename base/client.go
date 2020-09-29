package base

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/sirupsen/logrus"
)

type Base struct {
	ConfigPath   string `yaml:"configPath"`
	Organization string `yaml:"organization"`
	Username     string `yaml:"username"`
	ChannelID    string `yaml:"channelID"`
	ChainCodeID  string `yaml:"chainCode"`
}
type Client struct {
	// 生成fabSDK以及fabric go sdk中其他pkg使用的option context。
	fabricSDK *fabsdk.FabricSDK
	resClient *resmgmt.Client
	// 管理fabric网络中的成员关系
	mspClient *msp.Client
	// 实现Fabric账本的查询，查询区块、交易、配置等。
	ledgerClient *ledger.Client
	// 调用、查询Fabric链码，或者注册链码事件。
	channelClient *channel.Client
	// 配合channel模块来进行Fabric链码事件的注册和过滤。
	eventClient *event.Client
	base Base
}

func (c *Client) Init() {
	if len(c.base.ConfigPath) == 0 {
		return
	}
	if err := c.SetUp(); err != nil {
		logrus.Fatalf("failed to setup: %v", err)
	}
}
func (c *Client) SetUp() error {
	// 解析配置文件
	configProvider := config.FromFile(c.base.ConfigPath)
	// 通过配置文件创建fabric sdk go 入口实例
	fabricSDK, err := fabsdk.New(configProvider)
	if err != nil {
		return fmt.Errorf("failed to new fabricSDK %v", err)
	}
	c.fabricSDK = fabricSDK
	// 获取配置文件的用户名和组织
	clientProvider := c.fabricSDK.Context(
		fabsdk.WithUser(c.base.Username),
		fabsdk.WithOrg(c.base.Organization),
	)
	// 通过resmgmt.New创建fabric go sdk资源管理客户端
	resClient, err := resmgmt.New(clientProvider)
	if err != nil {
		return fmt.Errorf("failed to new resClient: %v", err)
	}
	c.resClient = resClient
	// 通过resmgmt.New创建fabric go sdk成员服务客户端
	if err != nil {
		return fmt.Errorf("failed to new mspClient: %v", err)
	}
	mspClient, err := msp.New(clientProvider)
	c.mspClient = mspClient
	// 通过channelID初始化ledger，channel，event客户端
	if c.base.ChannelID != "" {
		channelProvider := c.fabricSDK.ChannelContext(
			c.base.ChannelID,
			fabsdk.WithOrg(c.base.Organization),
			fabsdk.WithUser(c.base.Username),
		)
		ledgerClient, err := ledger.New(channelProvider)
		if err != nil {
			return fmt.Errorf("failed to new ledgerClient: %v", err)
		}
		c.ledgerClient = ledgerClient
		channelClient, err := channel.New(channelProvider)
		if err != nil {
			return fmt.Errorf("failed to new channelClient: %v", err)
		}
		c.channelClient = channelClient
		eventClient, err := event.New(channelProvider)
		if err != nil {
			return fmt.Errorf("failed to new eventClient: %v", err)
		}
		c.eventClient = eventClient
	}
	return nil
}
