.PHONY: run build stop clean

run:
	docker-compose up --build -d

build:
	docker-compose build

stop:
	docker-compose down

clean:
	docker-compose down -v
	docker system prune -f

logs:
	docker-compose logs -f

test:
	cd backend && go test ./...
	cd frontend && npm test 