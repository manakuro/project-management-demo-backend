# Asana clone app for backend
An Asana clone application for portfolio apps, built with real-time communication with WebSockets and seamless updates with Optimistic UI.

- [View project](https://project-management-demo.manatoworks.me/)
- [GraphQL Playground](https://project-management-demo.ebad78r3fqm9m.ap-northeast-1.cs.amazonlightsail.com/playground)
- [View front-end](https://github.com/manakuro/project-management-demo-frontend)

## Run Docker

```
$ cd docker
$ docker compose up
```

## Installation

```
$ make install
```

## Set up database

```
$ make setup_db
```

```
$ make migrate_schema
```

```
$ make seed
```

## Start server

```
$ make start
```
