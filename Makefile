run:
	go run ./cmd/game

build:
	go build -ldflags='-s' -o ./bin/gopher-pacman ./cmd/game