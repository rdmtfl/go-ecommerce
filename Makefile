up:
	@echo "Starting docker images ..."
	docker compose up -d
	@echo "... done!"

down:
	@echo "Stopping docker images ..."
	docker compose down
	@echo "... done!"

build:
	@echo "Stopping docker images (if running ...)"
	docker compose down
	@echo "Building and starting docker images ..."
	docker compose up --build -d
	@echo "... done!"