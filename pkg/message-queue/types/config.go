package types

type MessageQueueConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type QueueConfig struct {
	QueueName        string `json:"queue_name"`
	Durable          bool   `json:"durable"`
	DeleteWhenUnused bool   `json:"delete_when_unused"`
	Exclusive        bool   `json:"exclusive"`
	NoWait           bool   `json:"no_wait"`
}

type ConsumeConfig struct {
	QueueName string `json:"queue_name"`
	Consumer string `json:"consumer"`
	AutoAck bool `json:"auto_act"`
	Exclusive bool `json:"exclusive"`
	NoLocal bool `json:"no_local"`
	NoWait bool `json:"no_wait"`
}
