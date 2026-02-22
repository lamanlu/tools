TargetPath := .
AppName := tools

all:build

build:
	@echo "Building..."
	@cd $(CURDIR) && rm -f $(TargetPath)/$(AppName) && go build -o $(TargetPath)/$(AppName) ./main.go
	@echo "Done"
