package other

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

type ProductAggregator struct {
	concurrency int
	client      *http.Client
}

type Result struct {
	url   string
	price float32
	err   error
}

type Option func(*ProductAggregator) error

func WithConcurrency(concurrency int) Option {

	return func(aggregator *ProductAggregator) error {
		if concurrency <= 0 {
			return fmt.Errorf("concurrency must be positive")
		}
		aggregator.concurrency = concurrency
		return nil
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(aggregator *ProductAggregator) error {
		if timeout <= 0 {
			return fmt.Errorf("timeout must be positive")
		}
		aggregator.client.Timeout = timeout
		return nil
	}
}

func NewProductAggregator(options ...Option) (*ProductAggregator, error) {
	defaultAgg := &ProductAggregator{
		concurrency: runtime.GOMAXPROCS(0),
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
	for _, option := range options {
		if err := option(defaultAgg); err != nil {
			return nil, err
		}
	}
	return defaultAgg, nil
}

func (p *ProductAggregator) Aggregate(productUrls []string) []Result {
	workCh := make(chan string, p.concurrency)
	resultsCh := make(chan Result, len(productUrls))

	var wg sync.WaitGroup
	wg.Add(p.concurrency)

	for i := 0; i < p.concurrency; i++ {
		go func() {
			for url := range workCh {
				resultsCh <- p.query(url)
			}
			wg.Done()
		}()
	}

	for _, url := range productUrls {
		workCh <- url
	}
	close(workCh)
	wg.Wait()
	close(resultsCh)

	results := make([]Result, 0, len(productUrls))
	for result := range resultsCh {
		results = append(results, result)
	}
	return results
}

func (p *ProductAggregator) query(productUrl string) Result {
	result := Result{url: productUrl}
	rs, err := p.client.Get(productUrl)
	if err != nil {
		result.err = fmt.Errorf("error fetching product data: %w", err)
		return result
	}
	defer func() {
		if err := rs.Body.Close(); err != nil {
			//this could either be logged or injected to *Result (return type would have to be pointer)
			log.Printf("error closing response body: %v", err)
		}
	}()

	rsJson := new(priceResponse)
	err = json.NewDecoder(rs.Body).Decode(rsJson)
	if err != nil {
		result.err = fmt.Errorf("error decoding product data: %w", err)
		return result
	}
	result.price = rsJson.Price
	return result
}

type priceResponse struct {
	Price float32 `json:"price"`
}
