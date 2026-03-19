# Coffee_FS_App - System Flow

## Request Flow

Browser
-> React App (coffee-client/src/App.jsx)
-> Vite proxy (coffee-client/vite.config.js)
-> Go Router (coffee-server/router/router.go)
-> Controller Handler (coffee-server/controllers/controller.go)
-> Service Method (coffee-server/services/coffee.go)
-> DB Connection (coffee-server/db/db.go)
-> PostgreSQL tables (coffee-server/migrations/\*.sql)

## Endpoint Map (Current)

- `GET /health`
- `GET /api/coffees/all`
- `POST /api/coffees/create`
- `GET /api/coffees/id/{id}`
- `GET /api/coffees/name/{name}`
- `GET /api/coffees/query?region=...`
- `GET /api/coffees/query/{region}`
- `GET /api/coffees/price?price=...`
- `PUT /api/coffees/{id}`
- `DELETE /api/coffees/{id}`
