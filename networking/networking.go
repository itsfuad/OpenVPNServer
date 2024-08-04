// networking/networking.go
package networking

import (
    "crypto/tls"
    "fmt"
    "io"
    "log"
    "net"
    "vpn-server/authentication"
    "vpn-server/config"
    "vpn-server/encryption"
)

func Initialize(cfg *config.Config) error {
    address := net.JoinHostPort(cfg.ServerAddress, fmt.Sprint(cfg.Port))
    listener, err := tls.Listen("tcp", address, encryption.GetTLSConfig())
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

    tlsConn, ok := conn.(*tls.Conn)
    if !ok {
        log.Println("Failed to cast connection to TLS")
        return
    }

    state := tlsConn.ConnectionState()
    if len(state.PeerCertificates) == 0 {
        log.Println("No client certificate provided")
        return
    }

    clientCert := state.PeerCertificates[0]
    username := clientCert.Subject.CommonName

    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        log.Printf("Error reading password: %v", err)
        return
    }
    password := string(buffer[:n])

    if err := authentication.Authenticate(username, password); err != nil {
        log.Printf("Authentication failed for user %s: %v", username, err)
        return
    }

    log.Printf("User %s authenticated successfully", username)

    // Handle the VPN connection logic here
    for {
        n, err := conn.Read(buffer)
        if err != nil {
            if err != io.EOF {
                log.Printf("Error reading from connection: %v", err)
            }
            break
        }

        // Echo back the received data
        _, err = conn.Write(buffer[:n])
        if err != nil {
            log.Printf("Error writing to connection: %v", err)
            break
        }
    }
}
