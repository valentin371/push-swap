# push-swap

Deux programmes en Go (bibliothèque standard uniquement) :
- push-swap : affiche les instructions qui trient la pile a
- checker : lit les instructions sur l'entrée standard et affiche OK ou KO

## Compiler
go build -o push-swap ./cmd/push-swap
go build -o checker ./cmd/checker

## Utiliser
ARG="4 67 3 87 23"; ./push-swap "$ARG" | wc -l
ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"

## Tests
go test ./...
go run ./cmd/bench 100 20