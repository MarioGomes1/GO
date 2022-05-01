shell:!/bin/bash
# psql -U postgres -d go_movies -c "SELECT id  FROM movies ";
psql -c 'DROP DATABASE IF EXISTS  users;'
psql -c 'CREATE DATABASE users;'



migrate -path db/migration -database "postgresql://postgres:@localhost:5432/users?sslmode=disable" -verbose up 