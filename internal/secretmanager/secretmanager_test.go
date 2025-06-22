package secretmanager

import "testing"

func TestSecretManager(t *testing.T) {
	k := "SMTP_TO"
	v := Getenv(k)
	if v == "" {
		t.Fatal("secret response value blank")
	}
}
