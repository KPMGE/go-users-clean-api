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
