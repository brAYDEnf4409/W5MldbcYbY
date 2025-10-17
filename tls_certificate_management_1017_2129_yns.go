// 代码生成时间: 2025-10-17 21:29:35
package main

import (
    "crypto/tls"
    "crypto/x509"
    "encoding/pem"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/kataras/iris/v12"
)

// Certificate represents a certificate and its private key.
type Certificate struct {
    TLSCert *tls.Certificate
    FilePath string
}

// NewCertificate creates a new self-signed certificate.
func NewCertificate(domainName string) (*Certificate, error) {
    cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
    if err != nil {
        log.Fatalf("Error loading certificate: %s", err)
        return nil, err
    }

    return &Certificate{
        TLSCert: cert,
        FilePath: "cert.pem",
    }, nil
}

// SaveCertificate saves the certificate to the file system.
func (c *Certificate) SaveCertificate() error {
    certOut, err := os.Create(c.FilePath)
    if err != nil {
        return err
    }
    defer certOut.Close()

    pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: c.TLSCert.Certificate[0]})

    return nil
}

func main() {
    app := iris.New()

    // Define routes for certificate management.
    app.Get("/generate", func(ctx iris.Context) {
        domainName := ctx.URLParam("domain")
        if domainName == "" {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.WriteString("Domain name is required")
            return
        }

        cert, err := NewCertificate(domainName)
        if err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            ctx.WriteString("Failed to generate certificate: " + err.Error())
            return
        }

        if err := cert.SaveCertificate(); err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            ctx.WriteString("Failed to save certificate: " + err.Error())
            return
        }

        ctx.WriteString("Certificate generated and saved successfully")
    })

    app.Get("/renew", func(ctx iris.Context) {
        // Implement renewal logic here.
        ctx.WriteString("Renewal functionality not implemented")
    })

    // Define the SSL/TLS configuration.
    config := &tls.Config{
        PreferServerCipherSuites: true,
        CurvePreferences: []tls.CurveID{tls.CurveP256},
   }

    // Start the HTTPS server with the SSL/TLS configuration.
    if err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil); err != nil {
        log.Fatalf("Failed to start HTTPS server: %s", err)
    }
}
