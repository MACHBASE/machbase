package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/machbase/neo-grpc/machrpc"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	serverAddr := "127.0.0.1:5655"
	serverCert := filepath.Join(homeDir, ".config", "machbase", "cert", "machbase_cert.pem")

	// This example substitute server's key & cert for the client's key, cert.
	// It is just for the briefness of sample code
	// Client applicates **SHOULD** issue a certificate for each one.
	// Please refer to the "API Authentication" section of the documents.
	clientKey := filepath.Join(homeDir, ".config", "machbase", "cert", "machbase_key.pem")
	clientCert := filepath.Join(homeDir, ".config", "machbase", "cert", "machbase_cert.pem")

	cli := machrpc.NewClient(
		machrpc.WithServer(serverAddr),
		machrpc.WithCertificate(clientKey, clientCert, serverCert))
	if err := cli.Connect(); err != nil {
		panic(err)
	}
	defer cli.Disconnect()

	var count int

	sqlText := `select count(*) from example`
	row := cli.QueryRow(sqlText)
	if err := row.Scan(&count); err != nil {
		panic(err)
	}
	fmt.Println("count=", count)
}
