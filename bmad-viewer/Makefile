.PHONY: dev build run clean

dev:
	@echo "Starting dev servers..."
	cd web && npm run dev &
	cd server && go run . &
	wait

build:
	cd web && npm run build
	rm -rf server/dist
	cp -r web/dist server/dist
	cd server && go build -o ../bmad-viewer .

run: build
	./bmad-viewer

clean:
	rm -rf web/dist server/dist bmad-viewer
