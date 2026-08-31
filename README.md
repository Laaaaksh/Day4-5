# Day4-5

A basic REST API backend written in Go, built as a practice exercise around
July 2022 (commit history spans 2022-07-11 to 2022-07-18).

## What it is

A small "product + transaction" CRUD API: create/read/update/delete products,
and place transactions against them with basic username/password
authentication on the write endpoints. It reads like a learning exercise
working through Gin, GORM, and MySQL for the first time rather than a
production service.

## Stack

- Go 1.18
- [Gin](https://github.com/gin-gonic/gin) for HTTP routing
- [GORM](https://github.com/jinzhu/gorm) (the older v1 package) with a MySQL driver
- MySQL for storage, auto-migrated on startup

## Endpoints

```
/product-api/product        GET, POST
/product-api/product/:id    GET, PUT, DELETE
/transaction-api/transaction        GET, POST
/transaction-api/transaction/:id    GET, PUT, DELETE
```

Write endpoints (`POST`/`PUT` on products, and all of the transaction routes)
expect HTTP Basic Auth; credentials are checked against a `user` table with
a plaintext password comparison.

## Notable implementation details

- Database credentials (host, user, password) are hardcoded in
  `config/Database.go` rather than read from the environment. Don't reuse
  that file as-is — treat it as a placeholder to replace with your own
  config if you build on this.
- Package casing is inconsistent between the `config` directory and the
  `Day4-5/Config` import path used throughout the code, which trips Go's
  case-insensitive import check.

## Status

Written in July 2022 while learning Go/Gin/GORM. Not maintained, and **does
not currently build**: `go build ./...` fails with
`case-insensitive import collision: "Day4-5/config" and "Day4-5/Config"`
because the source files import `Day4-5/Config` (capital C) while the
directory on disk is `config` (lowercase). Fixing that casing mismatch is
the first thing you'd need to do to run it locally against a MySQL instance.
