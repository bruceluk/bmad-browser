.PHONY: dev build clean

dev:
	@echo "Starting dev servers..."
	cd web && npm run dev &
	cd server && go run . &
	wait

build:
	cd web && npm run build
	cd server && go build -o ../bmad-viewer .

clean:
	rm -rf web/dist bmad-viewer
