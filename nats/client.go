// NATS SDK - Minimalista para MCPs Vertikon
package nats

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

// Client encapsula conexão NATS
type Client struct {
	conn *nats.Conn
}

// NewClient cria nova conexão NATS
func NewClient(url string) (*Client, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn}, nil
}

// Close fecha conexão
func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// Publish publica mensagem
func (c *Client) Publish(subject string, data interface{}) error {
	logger, _ := zap.NewProduction()
	
	msgBytes, err := json.Marshal(data)
	if err != nil {
		logger.Error("Failed to marshal message", zap.Error(err))
		return err
	}
	
	err = c.conn.Publish(subject, msgBytes)
	if err != nil {
		logger.Error("Failed to publish", zap.String("subject", subject), zap.Error(err))
		return err
	}
	
	logger.Info("Message published", zap.String("subject", subject))
	return nil
}