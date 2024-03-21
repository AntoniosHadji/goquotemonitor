start:
	docker run --name quotemonitor \
		--env-file ./setup/env \
		--restart unless-stopped \
		--network=server \
		-p 8080:8080 \
		-d ghcr.io/antonioshadji/goquotemonitor

postgres:
	docker run --name postgres15-server \
		--mount type=volume,src=data-postgres15,dst=/var/lib/postgresql/data \
		--mount type=volume,src=home,dst=/root \
		--restart unless-stopped \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-p 5432:5432 \
		-d postgres:15

grafana:
	docker run --name grafana-server \
		-p 3000:3000 \
		-d grafana/grafana

local-build:
	docker build -t quotes:ubuntu .

local-start:
	docker run --name quotemonitor \
		--env-file ./setup/env \
		--restart unless-stopped \
		--network=server \
		-p 8080:8080 \
		-d quotes:ubuntu

show-errors:
	docker logs quotemonitor 2>&1 | rg Error

db-connect:
	docker exec -it postgres15-server psql -U postgres --db primetrust

db-connect-server:
	docker run -it --rm -e PGPASSWORD=password postgres:15 psql -U postgres -h 10.0.0.99 -d primetrust
