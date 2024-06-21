# Include all environment variables in .env.local
# Use - to ignore error when .env.local not found
-include .env.local
export

# Variables
DOCKER_COMPOSE=docker-compose --file build/docker-compose.yml --project-directory . -p ${PROJECT_NAME}

# Run
.PHONY: setup serve
setup: build-dev-image db ganache db-migrate

serve:
	@${DOCKER_COMPOSE} run --service-ports --rm app sh -c 'go run ./cmd/serverd'

deploy-contract:
	@${DOCKER_COMPOSE} run --rm app sh -c "go run ./cmd/deploycontract"

# Helper
.PHONY: build-dev-image update-vendor solc ormmodel api-test teardown
build-dev-image:
	docker build -f build/app.Dockerfile -t ${PROJECT_NAME}-app:latest .
	-docker images -q -f "dangling=true" | xargs docker rmi -f

update-vendor:
	@${DOCKER_COMPOSE} run --rm api sh -c 'go mod tidy && go mod vendor'

test:
	@${DOCKER_COMPOSE} run --rm app sh -c 'go test ./... -count=1'

teardown:
	@${DOCKER_COMPOSE} down

local-env:
	@cp .env.sample .env.local

generate:
	@${DOCKER_COMPOSE} run --rm app sh -c 'go generate ./...'

solc:
	${DOCKER_COMPOSE} run --rm solc sh -c "\
		solc --evm-version berlin --overwrite --abi data/contracts/*.sol -o build/contracts && \
		solc --evm-version berlin --overwrite --bin data/contracts/*.sol -o build/contracts"

ormmodel:
	@${DOCKER_COMPOSE} run --rm app sh -c "sqlboiler psql -c sqlboiler.yaml"

# Database
.PHONY: db db-migrate db-drop
db:
	@${DOCKER_COMPOSE} up -d db

db-migrate:
	@${DOCKER_COMPOSE} run --rm db-migrate sh -c 'migrate -path /migrations -database "$$DB_URL" up'

db-drop:
	@${DOCKER_COMPOSE} run --rm db-migrate sh -c 'migrate -path /migrations -database "$$DB_URL" drop'

# For test app
.PHONY: ganache
ganache:
	@${DOCKER_COMPOSE} up ganache -d
