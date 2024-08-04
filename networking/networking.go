// networking/networking.go
package networking

import (
	"fmt"
	"log"
	"net"
	"vpn-server/config"
)

func Initialize(cfg *config.Config) error {
    address := net.JoinHostPort(cfg.ServerAddress, fmt.Sprint(cfg.Port))
    listener, err := net.Listen("tcp", address)
    if err != nil {
        return err
    }

    go func() {
        for {
            conn, err := listener.Accept()
            if err != nil {
                log.Printf("Error accepting connection: %v", err)
                continue
            }

            go handleConnection(conn)
        }
    }()

    return nil
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    // Handle the VPN connection logic here
}
