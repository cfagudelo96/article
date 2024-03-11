buf.format:
	buf format -w

buf.mod:
	buf mod update proto

buf.lint: buf.format
	buf lint proto

buf.generate: buf.lint
	buf generate

docker.db:
	docker run -p 127.0.0.1:5432:5432/tcp --name postgres -e POSTGRES_PASSWORD=postgres -d postgres

sqlc.gen:
	sqlc -f core/data/sqlc.yaml generate
