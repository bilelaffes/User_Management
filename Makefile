run_server:
	go mod download
	go run main.go

docker_compose_up:
	docker-compose -f docker-compose.yml -p users_management up

docker_compose_down:
	docker-compose -f docker-compose.yml -p users_management down

test:
	go test -v -count=1 Users/create_test.go Users/login_test.go Users/read_test.go Users/update_test.go Users/delete_test.go