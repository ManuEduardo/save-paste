build:
	./tailwindcss -i src/views/css/styles.css -o src/public/styles.css
	@templ generate view
	@go build -o bin/fullstackgo main.go 

test:
	@go test -v ./...
	
run: build
	@./bin/fullstackgo

tailwind:
	@./tailwindcss -i src/views/css/styles.css -o src/public/styles.css --watch

templ:
	@templ generate -watch 

migration: # add migration name at the end (ex: make migration create-cars-table)
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run src/cmd/migrate/main.go up

migrate-down:
	@go run src/cmd/migrate/main.go down

#Por cosas de linux xd
run-air:
	@/home/manueduardo/go/bin/air

run-templ:
	@/home/manueduardo/go/bin/templ generate -watch