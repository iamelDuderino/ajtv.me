package secretmanager

import (
	"context"
	_ "embed"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"google.golang.org/api/option"
)

//go:embed "env.json"
var dotenv []byte

// Google Secret Manager
// Instead of hosting a .env file, a local env.json file should be present
// This will be a developers authorization for service account to connect to GCP Secret Manager

// To generate a token, first ensure the user has the role "Service Account Token Creator", then
// run the following command:
// gcloud auth application-default login --project <pid> --impersonate-service-account <service account>

func Getenv(secret string) string {
	var (
		s   string
		pid = "ajtvme"
		ctx = context.Background()
	)
	c, err := secretmanager.NewClient(ctx, option.WithCredentialsJSON(dotenv))
	if err != nil {
		fmt.Println(err)
		return s
	}
	defer c.Close()
	r, err := c.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: "projects/" + pid + "/secrets/" + secret + "/versions/latest",
	})
	if err != nil {
		fmt.Println(err)
		return s
	}
	s = string(r.Payload.Data)
	return s
}
