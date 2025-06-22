package secretmanager

import (
	"context"
	_ "embed"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

const (
	pid = "ajtvme"
)

func Getenv(secret string) string {
	ctx := context.Background()
	c, err := secretmanager.NewClient(ctx)
	if err != nil {
		return ""
	}
	defer c.Close()
	r, err := c.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: "projects/" + pid + "/secrets/" + secret + "/versions/latest",
	})
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(r.Payload.Data)
}
