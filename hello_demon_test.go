package hellodemon

import (
	"context"
	"testing"

	"github.com/beelzebub-labs/beelzebub/v3/pkg/plugin"
)

func TestHelloDemonExecute(t *testing.T) {
	got, err := (HelloDemon{}).Execute(context.Background(), plugin.CommandRequest{
		Command:  "whoami",
		Protocol: "ssh",
	})
	if err != nil {
		t.Fatalf("Execute returned error: %v", err)
	}

	want := `hello demon: received "whoami" over ssh`
	if got != want {
		t.Fatalf("Execute = %q, want %q", got, want)
	}
}
