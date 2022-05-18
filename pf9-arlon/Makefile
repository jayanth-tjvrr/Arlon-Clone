# vim:filetype=make:tabstop=4:shiftwidth=4:noexpandtab:

BUILD_DIR := $(REPO_ROOT)/build
GIT_SHORT_COMMIT := $(shell cd $(REPO_ROOT); git rev-parse --short HEAD)
DOCKERFILES := $(shell find $(REPO_ROOT) -mindepth 2 -maxdepth 2 -type f -name Dockerfile)
BUILD_NUMBER ?= 0
AWS_ACCOUNT_ID := 514845858982
AWS_REGION ?= us-west-1
CONTAINER_TAG ?= container-tag
DOCKER_REGISTRY_HOSTNAME := $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com
DOCKER_LOGOUT_CMD := docker logout $(DOCKER_REGISTRY_HOSTNAME)
CONTAINER_FULL_TAG ?= container-full-tag

