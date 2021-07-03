default:
	@echo "=============building Local API============="
	docker build -f price-alert/Dockerfile -t price-alert .

up: default
	@echo "=============starting api locally============="
	docker-compose up

logs:
	docker-compose logs -f

down:
	docker-compose down

clean: down
	@echo "=============cleaning up============="
	docker system prune -f
	docker volume prune -f