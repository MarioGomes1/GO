!/bin/bash
psql -U postgres -d go_movies -c "SELECT id  FROM movies ";
psql -c 'CREATE DATABASE routing;'

# psql -U postgres createdb simple_ban;
-migrate

# migrate -path db/migration -database "postgresql://postgres:@localhost:5432/routing?sslmode=disable" -verbose up 