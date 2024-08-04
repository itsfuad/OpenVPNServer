package main

import (
    "log"
    "vpn-server/config"
    "vpn-server/networking"
    "vpn-server/encryption"
    "vpn-server/authentication"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig("config.json")
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Initialize networking
    if err := networking.Initialize(cfg); err != nil {
        log.Fatalf("Error initializing networking: %v", err)
    }

    // Initialize encryption
    if err := encryption.Initialize(cfg); err != nil {
        log.Fatalf("Error initializing encryption: %v", err)
    }

    // Initialize authentication
    if err := authentication.Initialize(cfg); err != nil {
        log.Fatalf("Error initializing authentication: %v", err)
    }

    // Start the VPN server
    log.Println("VPN server started successfully")
    select {} // Block forever
}
