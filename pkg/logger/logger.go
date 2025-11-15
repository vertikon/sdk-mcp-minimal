// Logger SDK - Minimalista para MCPs Vertikon
package logger

import (
	"go.uber.org/zap"
)

// NewProduction cria um logger de produção
func NewProduction() (*zap.Logger, error) {
	return zap.NewProduction()
}

// NewDevelopment cria um logger de desenvolvimento
func NewDevelopment() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

// NewNop cria um logger no-op
func NewNop() *zap.Logger {
	return zap.NewNop()
}