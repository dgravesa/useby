package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/dgravesa/useby/pkg/useby"
)

func main() {
	var projectID string
	var username, password string

	flag.StringVar(&projectID, "projectID", "", "GCP project ID")
	flag.StringVar(&username, "username", "", "Name of user to create")
	flag.StringVar(&password, "password", "", "Password of user to create")
	flag.Parse()

	// validate command line arguments
	errs := []string{}
	if projectID == "" {
		errs = append(errs, "projectID is required")
	}
	if username == "" {
		errs = append(errs, "username is required")
	}
	if password == "" {
		errs = append(errs, "password is required")
	}

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	// initialize client
	client, err := useby.NewDatastoreClient(projectID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := client.PutUser(context.Background(), username, password); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("new user created successfully:", username)
	}
}
