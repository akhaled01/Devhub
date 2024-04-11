# About the `API` folder

This directory is where interfacing with the frontend happens. Here, parsing frontend queries, and executing them appropriately is done.

1. The `auth`  directory takes care of authentication
   1. Signup
   2. Login
2. The `chat`  directory takes care of websocket chats (MIGHT CHANGE)
3. The `posts` directory takes care of interfacting with posts
   1. Getting all posts
   2. Creating posts
   3. filtering and sorting posts based on different requirements
   4. Dealing with likes and comments

For each folder, there is a `routes.go` file, this is where routes are registered for the server

The `server` directory has the main server struct, which contains the main `mux` router and the listening address (post 7000 in this case). It also holds the `Boot` method and the `shutdown` method that handles graceful shutdown
