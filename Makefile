TargetPath := .
AppName := tools

all:build

build:
	@echo "Building..."
	rm -f $(TargetPath)/$(AppName)
	go build -o $(TargetPath)/$(AppName) ./main.go
	@echo "Done"