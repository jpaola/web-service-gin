# web-service-gin
Golang Gin Framework CRUD RestAPI with MySQL for music records.

# Installation
This application runs on Go version 1.18.1. You can install Go directly or by using Homebrew.

Using Homebrew:

```
brew install go
```

For additional details see https://formulae.brew.sh/formula/go.

Getting started with Go: https://go.dev/doc/tutorial/getting-started.

For information on how to download and install Go without using Homebrew see https://go.dev/doc/install.

# Getting Started
To run the application simply run the following command:

```
go run .
```

# MySQL Database Setup
Make sure to setup a database with your own credentials following the format provided within the database directory.

Initial queries to get you started:

```
CREATE TABLE albums ( albumID int NOT NULL, title varchar(64) NOT NULL, artist varchar(64) NOT NULL, price double(4,2), PRIMARY KEY (albumID) );

Insert into albums (albumID, title, artist, price) VALUES (1, 'Blue Train', 'John Coltrane', 56.99);
Insert into albums (albumID, title, artist, price) VALUES (2, 'Jeru', 'Gerry Mulligan', 17.99);
Insert into albums (albumID, title, artist, price) VALUES (3, 'Sarah Vaughan and Clifford Brown', 'Sarah Vaughan', 33.99);
```
The table should have albumID, title, artist, and price where albumId is considered the primary key.
