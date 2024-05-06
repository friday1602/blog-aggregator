run: build
build:
	@go build -o out && ./out
	
up:
	goose -dir ./sql/schema postgres $(DATABASE_URI) up

down:
	@goose -dir ./sql/schema postgres $(DATABASE_URI) down

connectDB:
	@psql $(DATABASE_URI)

statusDB:
	@goose -dir ./sql/schema postgres $(DATABASE_URI) status