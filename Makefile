docker:
	docker build -t armory/demo-api:latest .

push:
	docker push armory/demo-api:latest