package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

func search(ctx context.Context, word string) (string, error) {
	if word == "Go" {
		return "", errors.New("error:Go")

	}
	return fmt.Sprintf("result:%s", word), nil
}
func coSearch(ctx context.Context, words []string) ([]string, error) {
	var (
		wg      = sync.WaitGroup{}
		results = make([]string, len(words))
		err     error
	)
	for i, word := range words {
		wg.Add(1)
		go func(word string, i int) {
			defer wg.Done()
			result, e := search(ctx, word)
			if e != nil {
				return
			}
			results[i] = result
		}(word, i)
	}
	wg.Wait()
	return results, err
}
func main() {
	words := []string{"Go", "Java", "PHP", "Rust", "JavaScript"}
	results, err := coSearch(context.Background(), words)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(results)
}
