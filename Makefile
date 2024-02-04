
#
#	gochat
#

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | grep -v fgrep | sed 's/\(.*\):.*## \(.*\)/\1 - \2/' | sort

## run-chat - runs the "chat" app
run-chat:
	go run cmd/chat/main.go

## lint - lint all of the go source
lint:
	golangci-lint run
	#go run -mod=mod github.com/golangci/golangci-lint/cmd/golangci-lint run

## setup-lint - install go install the linter
setup-lint:
	go install github.com/ysmood/golangci-lint@latest
#	brew install golangci-lint run --enable-all
