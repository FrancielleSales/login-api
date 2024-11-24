# login-api

Login API using Golang

## Creating the development environment

### Database

**With PostgreSQL installed, access it:**

```sh
psql -u postgres psql
```

**Create a user to manage the database:**

```sh
CREATE USER admin WITH PASSWORD '321';
```

**Create the database:**

```sh
CREATE DATABASE db_login;
```

**Grant permissions to the user:**

```sh
GRANT ALL PRIVILEGES ON DATABASE db_login TO admin;
```

To exit the shell, use: \q
