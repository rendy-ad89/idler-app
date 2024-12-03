# Idler App
This is a mock web-app / idle game. The goal is to gather as many cash as we can by using machines to generate cash. This app is sort of inspired by https://steamcookieclicker.github.io

## Tech Stack
### Backend
- Golang with [Gin](https://gin-gonic.com)
- [SQLC](https://sqlc.dev)
### Frontend
- React
### Database
- PostgreSQL

## How to setup
- Clone this repo.
- Create a database named `idler_app_db` and run `idler_app_db.sql`. This sql script will create all the tables and insert some master datas.
- `cd` to `/backend` folder and run the backend app using `go run .`
- `cd` to `/frontend` folder and run `npm install`, followed by `npm run start`
