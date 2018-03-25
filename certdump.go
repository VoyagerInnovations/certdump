package main
  
import (
    "crypto/tls"
    "fmt"
    "encoding/pem"
    "os"
)

func main() {

    if len(os.Args) < 2 {
        argsProg := os.Args[0]
        fmt.Printf("Usage: %v [host:port]\n", argsProg)
        os.Exit(1)
    }

    uri := os.Args[1]

    conf := &tls.Config{
         //InsecureSkipVerify: true,
    }

    conn, err := tls.Dial("tcp", uri , conf)
    if err != nil {
        fmt.Print(err)
        return
    }
    defer conn.Close()

    certChain := conn.ConnectionState().PeerCertificates

    for i := range certChain {
        cert := certChain[i]
        pem := string(pem.EncodeToMemory(&pem.Block{ Type:  "CERTIFICATE", Bytes: cert.Raw }))
        fmt.Printf("\n\nCertificate No: %v\n\n", i);
        fmt.Printf("Issuer: %v\n", cert.Issuer)
        fmt.Printf("Subject: %v\n", cert.Subject.CommonName)
        fmt.Printf("Not Before: %v\n", cert.NotBefore)
        fmt.Printf("Not After: %v\n", cert.NotAfter)

        fmt.Print(string(pem))
    }
}
