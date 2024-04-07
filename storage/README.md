# About the Storage folder

This folder is where Interfacing with `sqlite` is done.

The `creators` directory has methods to initialize a new Database with the required tables according to our DB Schema. The `new_tables.go` file consolidates all the methods and creates them.

The `interfaces` directory is responsible for any interface between the database and the backend Each subdirectory contains methods to create and return .

The `init.go` file contains the `Init()` method that is called to initialize the Database before running the server.

> The sqlite driver used for this project is `github.com/mattn/go-sqlite3`. It might have some issues when launching with other operating systems, so make sure to use Docker instead!
