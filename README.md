# ks

`ks` is a simple utility for encoding and decoding kubernetes secrets.

![](assets/ks.png)

### why though

newline chars fool me once, shame on me
fool me twice, write a command line tool

```
a simple utility for base64 encoding secrets for k8s and copying them to the clipboard

Usage:
  ks [flags]
  ks [command]

Available Commands:
  decode      decode a base64 encoded secret
  encode      base64 encode a secret
  gui         run ks in interactive mode
  help        Help about any command

Flags:
  -f, --file string   target an existing secret yaml file -> ks -f <secret-filename>.yaml|.yml
  -h, --help          help for ks
  -s, --silent        no std output - clipboard only mode
  -V, --verbose       verbose output

Use "ks [command] --help" for more information about a command.
```

## todo
- [ ] environment variable check on input targets