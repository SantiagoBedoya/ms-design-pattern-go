package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

type Microservice struct {
	Name       string
	URL        string
	Circuit    *gobreaker.CircuitBreaker
	HTTPClient *http.Client
}

func NewMicroservice(name, URL string) *Microservice {
	st := gobreaker.Settings{
		Name:    name,
		Timeout: 5 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.6
		},
	}

	cb := gobreaker.NewCircuitBreaker(st)

	return &Microservice{
		Name:       name,
		URL:        URL,
		Circuit:    cb,
		HTTPClient: &http.Client{},
	}
}

func (m *Microservice) Call() (string, error) {
	var resp *http.Response
	var err error

	result, err := m.Circuit.Execute(func() (interface{}, error) {
		req, err := http.NewRequest("GET", m.URL, nil)
		if err != nil {
			return nil, err
		}

		resp, err = m.HTTPClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexepected status code: %d", resp.StatusCode)
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return string(data), nil
	})

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Success: %s", result), nil
}

func main() {
	ms1 := NewMicroservice("Service 1", "http://localhost:8081")
	ms2 := NewMicroservice("Service 2", "http://localhost:8082")

	for i := 0; i < 10; i++ {
		result, err := ms1.Call()
		if err != nil {
			log.Printf("[%s] Error: %v", ms1.Name, err)
		} else {
			log.Printf("[%s] Result: %s", ms1.Name, result)
		}

		result, err = ms2.Call()
		if err != nil {
			log.Printf("[%s] Error: %v", ms2.Name, err)
		} else {
			log.Printf("[%s] Result: %s", ms2.Name, result)
		}

		time.Sleep(time.Second)

	}
}
