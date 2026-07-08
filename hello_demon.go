package hellodemon

import (
	"context"
	"fmt"
	"strings"

	"github.com/beelzebub-labs/beelzebub/v3/pkg/plugin"
)

func init() {
	plugin.Register(&HelloDemon{})
}

type HelloDemon struct{}

func (HelloDemon) Metadata() plugin.Metadata {
	return plugin.Metadata{
		Name:        "hello-demon",
		Version:     "0.1.0",
		Author:      "Beelzebub Labs",
		Description: "Template command plugin that verifies external plugin loading.",
	}
}

func (HelloDemon) Execute(_ context.Context, req plugin.CommandRequest) (string, error) {
	command := strings.TrimSpace(req.Command)
	if command == "" {
		command = "silence"
	}

	protocol := strings.TrimSpace(req.Protocol)
	if protocol == "" {
		protocol = "unknown"
	}

	return fmt.Sprintf("hello demon: received %q over %s", command, protocol), nil
}
