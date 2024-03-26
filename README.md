# Healthcare web system (WIP)

[![wakatime](https://wakatime.com/badge/user/81fc6a55-00e9-44a0-af36-d8740dfd90ce/project/018c21ae-4195-4196-860e-94f4ef5ad6a8.svg)](https://wakatime.com/badge/user/81fc6a55-00e9-44a0-af36-d8740dfd90ce/project/018c21ae-4195-4196-860e-94f4ef5ad6a8)

## Requirements

1.  Package manager: [npm](https://www.npmjs.com/), [yarn](https://yarnpkg.com/), [bun](https://bun.sh/), [pnpm](https://pnpm.io/)
2.  [Go](https://go.dev/) (minimal 1.21.0)
3.  Database provider: [mariaDB](https://mariadb.org/), [mysql](https://www.mysql.com/), [postgres](https://www.postgresql.org/), [sql Server](https://www.microsoft.com/en-us/sql-server/sql-server-downloads)

## Development installation

1.  Clone the package from the [github repo](https://github.com/Thrustgamers/Healthcare)
    
2.  Create a dotenv file in the *server folder* with the following structure
    
    |     |     |     |
    | --- | --- | --- |
    | COOKIE_ENCRYPT | string | A random 32 character string used to encrypt cookies |
    | DB_HOST | string | The ip address of the database |
    | DB_USER | string | The database priviliged user |
    | DB_PASS | string | The user password |
    | DB_NAME | string | The database name |
    | DB_PORT | string | The database port |
    | DB_DRIVER | string | The database provider (Postgres, mysql, sqlServer) |
    
3.  Run the install command on your package manager in the *web folder*
    
4.  Run the `go mod download` in the *server folder*
    
5.  The server can be started using `air` and the front end can be started using the dev command using your package manager
    

## Feature request

- If you would like to see a feature added or want to have a converstation about it to see what is possible, don't hesitate to make a feature request! All input is appreciated
