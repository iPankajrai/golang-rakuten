# Contacts

- Create a Contact 
- Read a Contact
- Update a Contact
- Delete a Contact 

- flags
- os
- log
- gin-gonic (Http package)
- GORM (Go ORM)
- Postgres
- Docker

## Points
    - There is no need any webserver to deploy go web applications
    - Go web applications are selfhosted

-- docker for database

```
docker run -d --name pg2 -p 65432:5432  -e POSTGRES_PASSWORD=admin123 -e POSTGRES_USER=admin -e POSTGRES_DB=contactsdb postgres
```

-- docker for postgres admin

```
docker run -d --name pgadmin -p 5050:80  -e PGADMIN_DEFAULT_PASSWORD=root -e PGADMIN_DEFAULT_EMAIL=admin@admin.com dpage/pgadmin4
```
- open brower http://localhost:5050
- it askes for email and password: admin@admin.com, root respectively

- to obtain ip address of your postgress 

```
docker inspect pg2
```

- use that ip address for creating a new connection in pgadmin. The port is 5432 not 65432 because it connects internally from docker not externally