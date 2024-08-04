package authentication

import (
    "errors"
    "sync"
    "vpn-server/config"
)

var users = map[string]string{
    "user1": "password1",
    "user2": "password2",
}

var mu sync.Mutex

func Initialize(cfg *config.Config) error {
    // Load users from a file or database if necessary
    return nil
}

func Authenticate(username, password string) error {
    mu.Lock()
    defer mu.Unlock()

    if pass, ok := users[username]; ok {
        if pass == password {
            return nil
        }
    }
    return errors.New("authentication failed")
}
