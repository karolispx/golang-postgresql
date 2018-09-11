# GoLang and PostgreSQL. Very simple examples for inserting, selecting and deleting from PostgreSQL DB.

## Installing PostgreSQL on your machine/server should be straight forward. You should be able to find tutorials on google/youtube. 
### I have installed it locally on my macbook.

## Install PostgreSQL Client on your machine to make your life easier, I'm using Postico on a macOS.

You will need to get the following package to make it work:
* `go get -u github.com/lib/pq`

## Books table:
```
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    bookid character varying(50) NOT NULL,
    bookname character varying(255) NOT NULL
);
```
