# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# CURL

token:
	curl -il -X GET http://127.0.0.1:4000/token/MyKID

token-stg:
	curl -il -X GET https://staging-class0624-8652.encr.app/token/MyKID

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
