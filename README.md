# golang-core

Common code used in Go projects

Usage:
```
go get github.com/unidev-platform/golang-core
```

Important directories:
  * `xfiles` - file related utilities
  * `xcollection` - collections related functions
  * `xstring` - strings operations


## Examples

List files from directory matching extension:
```
	list, err := xfiles.Find("/tmp/domains", ".txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range list {
		log.Printf("%v\n", file)
	}

```

# License

This code is released under the MIT License. See LICENSE.
