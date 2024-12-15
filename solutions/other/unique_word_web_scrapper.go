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
	jobs := make(chan string, len(urls))
	results := make(chan result, len(urls))

	var wg sync.WaitGroup
	wg.Add(s.concurrencyLevel)
	for i := 0; i < s.concurrencyLevel; i++ {
		go func() {
			defer wg.Done()
			for url := range jobs {
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
		}()
	}

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

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
