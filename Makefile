up:
	docker-compose up
down:
	docker-compose down
migrate-up:
	migrate -database 'mysql://root:root@tcp(127.0.0.1:33306)/tech_story' -path ./mysql/migrations up
migrate-down:
	migrate -database 'mysql://root:root@tcp(127.0.0.1:33306)/tech_story' -path ./mysql/migrations down
