!/bin/bash
psql -U postgres -d go_movies -c "SELECT id  FROM movies ";
psql -c 'CREATE DATABASE routing;'

# psql -U postgres createdb simple_ban;
-migrate

# migrate -path db/migration -database "removed://removed:@removed:removed/routing?sslmode=disable" -verbose up 
