# Api Documentation


### Environment variables
You are supposed to set up a __.env__ file at the root of this project. That file's gonna be used 
to configure your database among other stuff. Here it is an example of such a file. Don't forget to change
those default configurations.

```bash
POSTGRES_HOST="localhost"
POSTGRES_PASSWORD="root"
POSTGRES_PORT=5432
POSTGRES_USER="postgres"
POSTGRES_PASSWORD="root"
POSTGRES_DB_NAME="users"
```

### Postgres database
I've made a __docker-compose__ file for creating a simple postgres databse container for you. To use it, 
just make sure you've got docker and docker-compose installed on your machine. Then run the simple command: 

```bash
sudo docker-compose up -d
```

### Tables on the database
To create the tables inside your database, you can use the following queries: 

##### Users
```sql
CREATE TABLE users(
  id         UUID, 
  created_at TIMESTAMP(6),
  updated_at TIMESTAMP(6),
  name       VARCHAR(255),
  user_name  VARCHAR(255),
  email      VARCHAR(255)
);
```

##### Accounts
```sql
CREATE TABLE accounts(
  id         UUID, 
  created_at TIMESTAMP(6),
  updated_at TIMESTAMP(6),
  user_name  VARCHAR(255),
  email      VARCHAR(255),
  password   VARCHAR(255)
);
```

--- 


After a few seconds, your container should be up and running. The default database is __users__ and 
the default password, root. You can change them to whatever you want later!

### How to generate api docs.
You can easily generate your api docs. First of all, you're supposed to run the following command:

```bash
npx insomnia-documenter --config ./golang-api.json
```

Then, you can run a server with your brand new doc by using:
```bash
npx serve
```

Now it's as simple as opening a new tab on your favourite browser and access the link:
> http://localhost:3000
