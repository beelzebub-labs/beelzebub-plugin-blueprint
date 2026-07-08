# Beelzebub Plugin Blueprint

Minimal external plugin repository for Beelzebub.

It is both a template for new plugins and a local verifier for the plugin
installation flow introduced by `configurations/plugins.yaml`.

## Files

- `plugins.yaml`: plugin manifest read by `beelzebub plugin install`
- `go.mod`: standalone plugin module
- `hello_demon.go`: self-registering command plugin
- `hello_demon_test.go`: small behavior test

## Install From GitHub

From the Beelzebub repository:

```bash
beelzebub plugin install https://github.com/beelzebub-labs/beelzebub-plugin-blueprint --no-build
```

Or add it manually to `configurations/plugins.yaml`:

```yaml
plugins:
  - source: https://github.com/beelzebub-labs/beelzebub-plugin-blueprint
```

Then run:

```bash
beelzebub plugin install --no-build
beelzebub plugin list
```

The installed plugin should appear as `hello-demon`.

## Verify Local Changes

For local development before pushing:

```bash
go test ./...
```

## Use As A Template

1. Rename the module in `go.mod`.
2. Update `plugins.yaml` so `module` matches `go.mod`.
3. Rename `HelloDemon` and update `Metadata`.
4. Keep the `init()` registration:

```go
func init() {
    plugin.Register(&YourPlugin{})
}
```
