docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres

migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up


/////////////////////////////////////////////////////////////////////////////////////////////////////////
Dirty database version 1. Fix and force version: solve:
migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" force <VERSION>
/////////////////////////////////////////////////////////////////////////////////////////////////////////


update schema_migrations set version='000001', dirty=false
