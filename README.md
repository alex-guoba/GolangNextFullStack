
# GolangNextFullStack - Full Stack App with Golang And Next.js


## Introduction 

This is a template for building full stack web applications with Golang and Next.js. It includes the following features:
- A simple CRUD API built with Gin framework (Golang)
- A simple React app built with Next.js
- Mysql database with migration support
- Dockerized development environment

GolangNextFullStack is an ideal project template suitable for developers who wish to leverage the high performance of Golang and the modern frontend capabilities of Next.js. Whether for individual projects or team collaboration, it allows for the quick launch and operation of a fully functional full-stack web application.

## Quick start

1. Clone this repo
2. From the project root directory, run
```sh
docker compose up -d
```
3. Open http://localhost:3000/


![BlogService](https://github.com/user-attachments/assets/404f6453-017c-4376-9f5c-99b904594abb)

## Architecture && TechStatck

![GolangNextFullStack](https://github.com/user-attachments/assets/97866ea0-83d3-4772-b29d-5fb7fbe359ef)

### Golang API

The API is built with Gin framework. It includes the following endpoints:
- `GET /api/v1/articles`: Get all articles
- `POST /api/v1/articles`: Create a new article
- `GET /api/v1/tags`: Get all tags
- `POST /api/v1/tags`: Create a new tag

We used [gin-clean-template](https://github.com/alex-guoba/gin-clean-template) as a base for the API.

### Next.js App

The app is built with Next.js( Next 15.1). It includes the following features:
- Two pages for displaying articles and tags.
- Some API routes for fetching articles and tags.
- Some components for displaying articles and tags list.

### Database

The database is MySQL. It includes the following tables:
- `articles`: Articles table with title, content and created_at fields
- `tags`: Tags table with name field
- `article_tags`: Many-to-many relationship table between articles and tags

## Development

### Running locally

To run the project locally, you need to have Docker installed on your machine. From the project root directory, run:
```sh
docker compose up
```
This will start the development environment including MySQL database, Go API server, and Next.js app. The API server will be available at http://localhost:8080 and the Next.js app will be available at http://localhost:3000.


## Tech stack

### Backend
- Golang
- [Gin framework](https://github.com/gin-gonic/gin): Fast and efficient HTTP web framework for Go
- [air](https://github.com/air-verse/air): Live reload for Go apps
- [Gorm](https://gorm.io/index.html): ORM library for Go
- Swagger: API documentation tool
- [golang-migrate](https://github.com/golang-migrate/migrate): Database migration tool for Go

### Frontend
- Next.js (Next 15.1): The React framework for building server-rendered applications
- React: A JavaScript library for building user interfaces
- [Tailwind CSS](https://tailwindcss.com/): A utility-first CSS framework for rapidly building custom designs
- [Shadcn/ui](https://ui.shadcn.com/): A collection of React UI components
- [T3 Environment](https://github.com/t3-oss/env): Environment variables for Next.js

