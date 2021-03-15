package stock

import (
	"context"
	"time"

	"github.com/gaius-qi/honk/internal/config"
)

type stockSupper interface {
	Get() (*Stock, error)
}

type stockContext struct {
	strategy stockSupper
}

type Stock struct {
	Name                 string
	Number               string
	OpeningPrice         string
	PreviousClosingPrice string
	CurrentPrice         string
	HighPrice            string
	LowPrice             string
	Date                 time.Time
}

func NewStockContext(ctx context.Context, platformType config.PlatformType, cfg *config.Config) stockSupper {
	s := new(stockContext)
	switch platformType {
	case config.SinaPlatformType:
		s.strategy = newSinaStock(cfg)
	default:
		s.strategy = newSinaStock(cfg)
	}
	return s
}

func (s stockContext) Get() (*Stock, error) {
	return s.strategy.Get()
}
