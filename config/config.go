package config

import (
    "encoding/json"
    "os"
)

type Config struct {
    ServerAddress string `json:"server_address"`
    Port          int    `json:"port"`
    CertFile      string `json:"cert_file"`
    KeyFile       string `json:"key_file"`
    CAFile        string `json:"ca_file"`
}

func LoadConfig(file string) (*Config, error) {
    f, err := os.Open(file)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    var cfg Config
    decoder := json.NewDecoder(f)
    if err := decoder.Decode(&cfg); err != nil {
        return nil, err
    }

    return &cfg, nil
}
