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

--- 

### Postgres database
I've made a __docker-compose__ file for creating a simple postgres databse container for you. To use it, 
just make sure you've got docker and docker-compose installed on your machine. Then run the simple command: 

```bash
sudo docker-compose up -d
```

--- 

### Tables on the database
In this project, i have used the *gorm* framework to deal with the migrations. So, you don't need to 
worry about it, it will create the tables automatically when you run the program.

After a few seconds, your container should be up and running. The default database is __users__ and 
the default password, root. You can change them to whatever you want later!

### How do i run it?
In order to run this project, make sure you have got golang properly installed on your machine, then run 
the command at the root of the project

```bash
go run ./src/main/main.go
```

--- 

### How to generate api route docs.
You can easily generate your api route docs. First of all, you're supposed to run the command below, make sure you've got
npx on your machine.

```bash
npx insomnia-documenter --config ./golang-api.json
```

Then, you can run a server with your brand new doc by using:
```bash
npx serve
```

Now it's as simple as opening a new tab on your favourite browser and access the link:
> http://localhost:3000
