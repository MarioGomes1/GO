shell:!/bin/bash
# psql -U postgres -d go_movies -c "SELECT id  FROM movies ";
psql -c 'DROP DATABASE IF EXISTS  users;'
psql -c 'CREATE DATABASE users;'


<<<<<<< HEAD

migrate -path db/migration -database "postgresql://postgres:@localhost:5432/users?sslmode=disable" -verbose up 
=======
# migrate -path db/migration -database "removed://removed:@removed:removed/routing?sslmode=disable" -verbose up 
>>>>>>> d797a80d6308565ddf785bc6ffd2e7c6601d5d56
