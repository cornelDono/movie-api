.PHONY: build sqlc clean

# Compile the Go code
build:
	air

# Run sqlc generate
sqlc:
	docker run --rm -v "/c/Users/User/Desktop/projects/GO/movie-api:/src" -w /src sqlc/sqlc generate

# Clean up binaries
clean:
	rm -rf bin/