package coinmarketcap

import (
	"fmt"
	"strings"
	"time"
)

type CryptocurrencyMapStore struct {
	client *CmcApiClient
	store map[string]CryptocurrencyMapEntryApiResult
}

func convertCryptocurrencyMap(cryptocurrencyMap CryptocurrencyMapApiResult) (map[string]CryptocurrencyMapEntryApiResult) {
	data := cryptocurrencyMap.Data

	entryCount := len(data)

	result := make(map[string]CryptocurrencyMapEntryApiResult, entryCount)

	for _, entry := range cryptocurrencyMap.Data {
		result[strings.ToUpper(entry.Symbol)] = entry
	}

	return result
}

func (s *CryptocurrencyMapStore) GetEntry(symbol string) (CryptocurrencyMapEntryApiResult) {
	return s.store[strings.ToUpper(symbol)]
}

func (s *CryptocurrencyMapStore) start() (error) {
	fmt.Print("CryptocurrencyMapStore starting...\n")
	err := s.update()

	if err != nil {
		return err
	}

	tick := time.Tick(4 * time.Hour)

	go func() {
		for range tick {
			go func() {
				fmt.Print("CryptocurrencyMapStore updating...\n")
				err := s.update()

				if err == nil {
					fmt.Printf("CryptocurrencyMapStore updated\n")
				} else {
					fmt.Printf("CryptocurrencyMapStore error while updating: %v\n", err)
				}
			}()
		}
	}()

	fmt.Printf("CryptocurrencyMapStore started\n")

	return nil
}

func (s *CryptocurrencyMapStore) update() (error) {
	cryptocurrencyMap, err := s.client.GetCryptocurrencyMap()

	if err != nil {
		return err
	}

	s.store = convertCryptocurrencyMap(cryptocurrencyMap)

	return nil
}

func NewCryptocurrencyMapStore() (*CryptocurrencyMapStore, error) {
	result := &CryptocurrencyMapStore{
		client: &CmcApiClient{},
	}

	err := result.start()

	if err != nil {
		return nil, err
	}

	return result, nil
}

