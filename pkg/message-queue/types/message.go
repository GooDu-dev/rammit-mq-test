package types

type Message interface {
	Metadata() MessageMetadata
}

type MessageMetadata struct {
	Exchange    string `json:"exchange"`
	RouteingKey string `json:"routing_key"`
	Mandatory   bool   `json:"mandatory"`
	Immediate   bool   `json:"immediate"`
}

type message struct {
	metadata MessageMetadata
}

func NewMessage(cfg MessageMetadata) Message {
	return &message{metadata: cfg}
}

func (m *message) Metadata() MessageMetadata {
	return m.metadata
}
