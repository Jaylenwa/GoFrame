package nsq

import (
	"GoFrame/infrastructure/config/pkg/util"
	"github.com/nsqio/go-nsq"

	"GoFrame/global"
)

type NsqClient interface {
	NsqProducer() (*nsq.Producer, error)
	NsqConsumer(topic string, ch string) (*nsq.Consumer, error)
	GetNsqProducerLink() string
	GetNsqConsumerLink() string
}

var _ NsqClient = (*nsqClient)(nil)

type nsqClient struct {
}

func NewNsqClient() *nsqClient {
	return &nsqClient{}
}

func (n *nsqClient) NsqProducer() (*nsq.Producer, error) {
	addr := n.GetNsqProducerLink()
	nsq, err := nsq.NewProducer(addr, nsq.NewConfig())
	if err != nil {
		return nil, err
	}

	return nsq, nil
}

func (n nsqClient) NsqConsumer(topic string, ch string) (n2 *nsq.Consumer, err error) {
	//conf := nsq.NewConfig()
	//conf.LookupdPollInterval = 15 * time.Second
	return nsq.NewConsumer(topic, ch, nsq.NewConfig())
}

func (n nsqClient) GetNsqProducerLink() string {
	return util.ParseHost(global.Gconfig.Nsq.NsqProducerHost) + ":" + global.Gconfig.Nsq.NsqProducerPort
}

func (n nsqClient) GetNsqConsumerLink() string {
	return util.ParseHost(global.Gconfig.Nsq.NsqSubscribeHost) + ":" + global.Gconfig.Nsq.NsqSubscribePort
}
