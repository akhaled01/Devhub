# About the `Storage` folder

This folder is where Interfacing with the `sqlite` database is done.

The `creators` directory has methods to initialize a new Database with the required tables according to our DB Schema. The `new_tables.go` file consolidates all the methods and creates them.

The `interfaces` directory is responsible for any interface between the database and the backend Each subdirectory contains methods to create and get different entities from the DB. Each function is documented well.

> Each directory in the `interfaces` directory has a `get.go` file to house methods tha concern with returning the entity (user for example) from the database and a `create.go` file that has methods to save entities to a database

The `init.go` file contains the `Init()` method that is called to initialize the Database before running the server.

> The sqlite driver used for this project is `github.com/mattn/go-sqlite3`. It might have some issues when launching with other operating systems, so make sure to use Docker instead!
