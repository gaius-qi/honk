package stock

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/gaius-qi/honk/internal/config"
)

const (
	defaultSinaHost     = "hq.sinajs.cn"
	defaultSinaProtocol = "https:"
	sinaTimeLayout      = "2006-01-02 15:04:05"
)

type sinaStock struct {
	number string
	index  config.IndexType
}

func newSinaStock(index config.IndexType, number string) sinaStock {
	return sinaStock{
		number: number,
		index:  index,
	}
}

func (s sinaStock) Get() (*Stock, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s//%s", defaultSinaProtocol, defaultSinaHost), nil)
	if err != nil {
		return nil, err
	}

	// Add query params
	q := req.URL.Query()
	q.Add("list", fmt.Sprintf("%s%s", s.index, s.number))
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Match useful data
	data := regexp.MustCompile(`[\\"\\,]+`).Split(string(body), -1)
	if len(data) < 35 {
		return nil, errors.New("Sina platform returns wrong data")
	}

	// Time parse
	t, err := time.Parse(sinaTimeLayout, fmt.Sprintf("%s %s", data[31], data[32]))
	if err != nil {
		return nil, err
	}

	return &Stock{
		Name:                 data[1],
		Number:               s.number,
		OpeningPrice:         data[2],
		PreviousClosingPrice: data[3],
		CurrentPrice:         data[4],
		HighPrice:            data[5],
		LowPrice:             data[6],
		Date:                 t,
	}, nil
}
