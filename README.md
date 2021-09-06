# ks

`ks` is a simple utility for encoding and decoding kubernetes secrets.

```
go get github.com/kathleenfrench/ks
```

## why though?

newline chars fool me once, shame on me...fool me twice, write a command line tool.

essentially, `ks` ensures any value copied to the command line for encoding/decoding has newline and/or added whitespace characters stripped. it then copies the result to your clipboard.

you can control how much information is written to `stdout` with the `--verbose` and `--silent` flags. `--silent` is pretty self-explanatory, in that nothing is written to `stdout`. 

the default behavior is for the encoded/decoded value to print to the terminal. a `--verbose` flag just writes additional info logs that confirm the result was copied to the clipboard, or - in the case of parsing k8s files - out-putting the `yaml`.

# usage

### encode

```
$ ks encode supersecretk8svalue
$ ks e supersecretk8svalue
$ ks encode $SOME_ENVIRONMENT_VARIABLE
``` 

### decode

```
$ ks decode c3VwZXJzZWNyZXRrOHN2YWx1ZQ==
$ ks d c3VwZXJzZWNyZXRrOHN2YWx1ZQ==
$ ks decode $SOME_ENVIRONMENT_VARIABLE
```

### interactive mode

in 'interactive mode', rather than directly provide any value to `ks`, you instead provide the filepath to a `yaml|yml` secret file. 

`ks` will then parse the data key/value pairs and you can select them from a prompt UI, from which you can opt to encode or decode a given value.

then you are prompted whether you want to just copy that response or copy it and also open the target file as a temporary file in your text editor of choice. upon saving and exiting, `ks` will overwrite the actual target file with your changes.


## copy only mode

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

## editor mode

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

saved changes to <file>!
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