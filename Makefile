IMAGE_NAME?=ghcr.io/rdimitrov/good-repo-go:latest
CR_USERNAME?=rdimitrov
# replace with your GitHub PAT, should have read/write access for packages
CR_PAT?=ghp_1234567890abcdefghij1234567890abcdefghij

.PHONY: login
login:
	@echo "Logging in to GitHub Container Registry"
	@echo "${CR_PAT}" | docker login ghcr.io -u $(CR_USERNAME) --password-stdin

.PHONY: build-image
build-image:
	@echo "Building a safe image..."
	docker build -t $(IMAGE_NAME) .

.PHONY: build-malicious-image
build-malicious-image:
	@echo "Building a malicious image..."
	@echo "// Maliciously altered on $$(date)" >> main.go
	docker build -t $(IMAGE_NAME) .


.PHONY: push-image
push-image:
	@echo "Pushing image..."
	docker push $(IMAGE_NAME)

.PHONY: keygen
keygen:
	@cosign generate-key-pair


.PHONY: sign-keypair
sign-keypair:
	@cosign sign $(IMAGE_NAME) --key cosign.key

.PHONY: sign-oidc
sign-oidc:
	@cosign sign $(IMAGE_NAME)
