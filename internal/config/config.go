package config

import (
	"encoding/json"
	"os"
	"sort"
)

type Config struct {
	PackSizes []int `json:"pack_sizes"`
}

func LoadConfig(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cfg.PackSizes))) // descending
	return cfg.PackSizes, nil
}
