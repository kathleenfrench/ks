# ks

`ks` is a simple utility for encoding and decoding kubernetes secrets.

```
go get github.com/kathleenfrench/ks
```

## why though?

newline chars fool me once, shame on me...fool me twice, write a command line tool.

###  what `ks` does

- encode or decode a terminal input and copy it to your clipboard, and trust that newline/whitespace chars have been stripped
- parse k8s secret data directly from a file into a simple terminal UI, and select whether to encode/decode by key
- option to open and edit k8s secret files after copying converted secret values
- ability to add new secrets via `ks` where `ks` handles the encoding/decoding for you. 

# usage

## flags

#### `--verbose`, `--silent`

you can control how much information is written to `stdout` with the `--verbose` and `--silent` flags. 

#### `--file`

the `--file` flag is a filepath relative to the current working directory if you are using `ks` with an existing `yaml` configuration.

## commands

### encode

```
$ ks encode supersecretk8svalue
$ ks e supersecretk8svalue
$ ks encode $SOME_ENVIRONMENT_VARIABLE
``` 

when you run `ks encode` with the `-f` flag and append a filename for an existing k8s secret config, `ks` will encode all of the secret values under `data` and output them to the terminal.

by default, this result is copied to the clipboard, but this behavior can be disabled by appending the `--nocopy` flag

```
$ ks encode -f example.yaml

apiVersion: v1
data:
  faux-secret-key: Wm1GclpYTmxZM0psZEE9PQ==
  top-secret-key: ZEc5d2MyVmpjbVYw
kind: Secret
metadata:
  name: secret-sa-sample
type: Opaque

result copied to clipboard!
```

### decode

```
$ ks decode c3VwZXJzZWNyZXRrOHN2YWx1ZQ==
$ ks d c3VwZXJzZWNyZXRrOHN2YWx1ZQ==
$ ks decode $SOME_ENVIRONMENT_VARIABLE
```


when you run `ks decode` with the `-f` flag and append a filename for an existing k8s secret config, `ks` will decode all of the secret values under `data` and output them to the terminal.

by default, this result is copied to the clipboard, but this behavior can be disabled by appending the `--nocopy` flag

```
$ ks decode -f example.yaml

apiVersion: v1
data:
  faux-secret-key: fakesecret
  top-secret-key: topsecret
kind: Secret
metadata:
  name: secret-sa-sample
type: Opaque

result copied to clipboard!
```

## file mode

when you provide a `-f` value (or filepath) with no subsequent arguments or subcommands, `ks` will parse the data key/value pairs and you can select them from a prompt UI, from which you can opt to encode or decode a given value.

after, you will be prompted whether you want to just copy that response or copy it and also open the target file as a temporary file in your text editor of choice. upon saving and exiting, `ks` will overwrite the actual target file with your changes.


### copy only mode

```
$ ks -f example.yaml

? select a key [Use arrows to move, type to filter]
> faux-secret-key
  top-secret-key

? select a key: top-secret-key

dG9wc2VjcmV0

? do you want to decode or encode this value? [Use arrows to move, type to filter]
> decode
  encode
  exit

? select one [Use arrows to move, type to filter]
> copy only
  copy & open target file
  exit

? select one: copy only

topsecret

```

### editor mode

```
$ ks -f example.yaml

? select a key [Use arrows to move, type to filter]
> faux-secret-key
  top-secret-key
  Add New Secret

? select a key: top-secret-key

dG9wc2VjcmV0

? do you want to decode or encode this value? [Use arrows to move, type to filter]
> decode
  encode
  exit

? select one [Use arrows to move, type to filter]
  copy only
> copy & open target file
  exit

? select an editor [Use arrows to move, type to filter]
> Vim
  Emacs
  Atom
  Sublime
  VS Code
  quit

? view/edit example.yaml [? for help] [Enter to launch editor]

<opens example.yaml tmp file in chosen editor>
<save file & quit>

saved any changes to <file>!
```

### add a new secret

```
$ ks -f example.yaml

? select a key [Use arrows to move, type to filter]
  faux-secret-key
  top-secret-key
> Add New Secret

? Provide a valid secret key: test-key
? Provide a secret value: ******
? Update secret value before saving? [Use arrows to move, type to filter]
  Decode
> Encode
  Save As Is
  Quit

<saves example.yaml with new secret>
```


# development

## local

to work on `ks` locally, there are a number of useful `make` targets.

` > make`

```
Local `make` Commands:

  build            compile the ks binary to the workspace's root build directory
  install          install the ks binary to /usr/local/bin
  clean            delete the build output directory
  lint             lint the go code
  test             run unit tests
  help             see available make commands

```

## CI

a `lint` and `build/test` action runs on PR branches using the declared workflows in `.github/workflows`