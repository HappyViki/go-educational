# go-educational
Learning the language Go with a project

Make more convenient:
```
export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
```

Install:
```
go install github.com/HappyViki/go-educational
```

Run:
```
go-educational
```

Test:
```
go test -v ./...
```

Remove cache:
```
go clean -modcache
```

[Go get started](https://golang.org/doc/code.html)
