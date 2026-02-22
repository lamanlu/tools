TargetPath := .
AppName := tools

all:build

build:
	@echo "Building..."
	@cd $(CURDIR) && rm -f $(TargetPath)/$(AppName) && go build -mod=readonly -trimpath -buildvcs=false -ldflags "-s -w" -o $(TargetPath)/$(AppName) ./main.go
	@echo "Done"
