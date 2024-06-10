# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# CURL

token:
	curl -il -X GET http://127.0.0.1:4000/token/MyKID

tst:
	curl -il -X POST http://127.0.0.1:4000/test -d '{"status": "BILL"}'

token-stg:
	curl -il -X GET https://staging-class0624-8652.encr.app/token/MyKID

tst-stg:
	curl -il -X POST https://staging-class0624-8652.encr.app/test -d '{"status": "BILL"}'

# ==============================================================================
# Help Stuff
# Q: How to pass values.

help:
	encore run -v --help

# ==============================================================================
# Manage Project

up:
	encore run -v --browser never

FIND_DAEMON = $(shell ps | grep 'encore daemon' | grep -v 'grep' | cut -c 1-5)
SET_DAEMON = $(eval DAEMON_ID=$(FIND_DAEMON))

FIND_APP = $(shell ps | grep 'encore_app_out' | grep -v 'grep' | cut -c 1-5)
SET_APP = $(eval APP_ID=$(FIND_APP))

down-daemon:
	$(SET_DAEMON)
	if [ -z "$(DAEMON_ID)" ]; then \
		echo "daemon not running"; \
    else \
		kill -SIGTERM $(DAEMON_ID); \
    fi

down-app:
	$(SET_APP)
	if [ -z "$(APP_ID)" ]; then \
		echo "app not running"; \
    else \
		kill -SIGTERM $(APP_ID); \
    fi

down: down-app down-daemon

statsviz:
	open -a "Google Chrome" http://127.0.0.1:4000/debug/statsviz

# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy

tidy:
	go mod tidy

deps-list:
	go list -m -u -mod=readonly all

deps-upgrade:
	go get -u -v ./...
	go mod tidy

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all
