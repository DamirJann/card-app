test:
	docker-compose up --build -d

	docker exec app ./app add -h
	docker exec app ./app list -h
	docker exec app ./app remove -h
	docker exec app ./app -v
	docker exec app ./app list
	docker exec app ./app add -name title -answer answer
	docker exec app ./app list
	docker exec app ./app remove -name title
	docker exec app ./app list
	docker exec app ./app remove -all
	docker exec app ./app list

	docker-compose down

build:
	 go build -o "app-${TAG}" -buildvcs=false .

deploy:
	cd ansible && make deploy -e TAG="${TAG}"
