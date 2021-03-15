package stock

import (
	"context"
	"reflect"
	"testing"

	"github.com/gaius-qi/honk/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewStockContext(t *testing.T) {
	assert := assert.New(t)

	index := config.ShangHaiIndexType
	number := "600000"
	s := NewStockContext(context.Background(), config.SinaPlatformType, &config.Config{
		Index:  index,
		Number: number,
	})

	assert.NotNil(s)
	assert.NotNil(reflect.ValueOf(s).Elem())
}
