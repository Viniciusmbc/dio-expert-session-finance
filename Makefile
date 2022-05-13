build-image:
	docker build -t viniciuscaceres/finance .

run-app:
	docker-compose -f ./.devops/app.yml up -d

