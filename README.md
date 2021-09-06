# ks

`ks` is a simple utility for encoding and decoding kubernetes secrets.

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
``` 

### decode

```
$ ks decode c3VwZXJzZWNyZXRrOHN2YWx1ZQ==

$ ks d c3VwZXJzZWNyZXRrOHN2YWx1ZQ==
```

### interactive mode

in 'interactive mode', rather than directly provide any value to `ks`, you instead provide the filepath to a `yaml|yml` secret file. `ks` will then parse the data key/value pairs and you can select them from a prompt UI, from which you can opt to encode or decode a given value and copy it to the keyboard.

```
$ ks -f example.yaml

? select a key [Use arrows to move, type to filter]
> faux-secret-key
  top-secret-key

? select a key: top-secret-key

dG9wc2VjcmV0

? select one [Use arrows to move, type to filter]
> decode & copy to clipboard
  encode & copy to clipboard
  exit

? select one: decode & copy to clipboard

topsecret

```

---

## todo
- [ ] environment variable check on input targets
- [ ] open `yaml` file in editor of choice