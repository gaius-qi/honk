package stock

import (
	"context"
	"time"

	"github.com/gaius-qi/honk/internal/config"
)

type StockSupper interface {
	Get() (*Stock, error)
}

type stockContext struct {
	strategy StockSupper
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

func NewStockContext(ctx context.Context, cfg *config.Config) StockSupper {
	s := new(stockContext)
	switch cfg.Platform {
	case config.SinaPlatformType:
		s.strategy = newSinaStock(cfg.Index, cfg.Number)
	default:
		s.strategy = newSinaStock(cfg.Index, cfg.Number)
	}
	return s
}

func (s stockContext) Get() (*Stock, error) {
	return s.strategy.Get()
}
