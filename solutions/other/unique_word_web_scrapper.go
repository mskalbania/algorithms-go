package other

import (
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Scrapper struct {
	client           http.Client
	concurrencyLevel int
}

func NewScrapper(timeout time.Duration, concurrencyLevel int) *Scrapper {
	return &Scrapper{
		client:           http.Client{Timeout: timeout},
		concurrencyLevel: concurrencyLevel,
	}
}

type result struct {
	url         string
	uniqueWords map[string]int
	err         error
}

func (s *Scrapper) Scrap(urls []string) map[string]int {
	jobs := toChanel(urls)
	results := make(chan result, len(urls))

	var wg sync.WaitGroup
	wg.Add(s.concurrencyLevel)
	for i := 0; i < s.concurrencyLevel; i++ {
		go s.processUrl(jobs, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	wordsCount := make(map[string]int)
	for result := range results {
		if result.err != nil {
			log.Printf("error processing %s: %s", result.url, result.err)
			continue
		}
		for word, count := range result.uniqueWords {
			wordsCount[word] += count
		}
	}
	return wordsCount
}

func toChanel(urls []string) chan string {
	ch := make(chan string, len(urls))
	for _, url := range urls {
		ch <- url
	}
	close(ch)
	return ch
}

func (s *Scrapper) processUrl(urls <-chan string, results chan<- result, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range urls {
		result := result{url: url}
		rs, err := s.client.Get(url)
		if err != nil {
			result.err = err
			results <- result
			continue
		}
		body, err := io.ReadAll(rs.Body)
		if err != nil {
			result.err = err
			results <- result
			continue
		}
		_ = rs.Body.Close()
		fields := strings.Fields(string(body))
		result.uniqueWords = make(map[string]int)
		for _, field := range fields {
			word := strings.Trim(field, ".,!?()[]{}\\\"':;")
			if word != "" {
				result.uniqueWords[word]++
			}
		}
		results <- result
	}
}
