# shiny-parakeet

# Requirements
## Technical Requirements
1. go 1.19
2. [golang-migrate](https://github.com/golang-migrate/migrate) for migrating database
3. Postgresql

## Other requirement
1. Please find example.env in root and add required values and rename it as app.env

# How to set up
## with Docker
1. Pull postgres image
```
docker pull postgres:12-alpine
```
2. Create container for database
```
make postgres
```
3. Create database in container
```
make createdb
```
4. Migrate database
```
make migrateup
```
5. Test application
```
make test
```
6. Run application
```
make server
```

## Without docker
1. Migrate database
```
make migrateup
```
2. Test application
```
make test
```
3. Run application
```
make server
```

# API Documentation 
Please find API Documentation in ./wiki folder