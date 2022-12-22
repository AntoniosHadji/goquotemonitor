build:
	docker build -t quotes:ubuntu .

run:
	docker run -it --env-file ./setup/env quotes:ubuntu


postgres:
	docker run --name postgres15-server \
		--mount type=volume,src=data-postgres15,dst=/var/lib/postgresql/data \
		--mount type=volume,src=home,dst=/root \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-p 5432:5432 \
		-d postgres:15

grafana:
	docker run --name grafana-server \
		-p 3000:3000 \
		-d grafana/grafana

