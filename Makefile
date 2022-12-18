build:
	docker build -t quotes:ubuntu .

run:
	docker run -it --env-file ./setup/env quotes:ubuntu
