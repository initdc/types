default:
  just --list --unsorted --justfile {{justfile()}}

fmt:
  go fmt ./...

vet:
  go vet ./...

test DIR="..." :
  go test -cover -v ./{{ DIR }}

bench:
  go test -bench=. -benchmem -cpu='1,2,4'
