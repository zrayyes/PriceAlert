default:
	@echo "=============building Local API============="
	docker build -f price-alert/Dockerfile -t price-alert .

up:
	@echo "=============starting api locally============="
	docker-compose up --build

logs:
	docker-compose logs -f

down:
	docker-compose down

clean: down
	@echo "=============cleaning up============="
	docker system prune -f
	docker volume prune -f

test:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down 