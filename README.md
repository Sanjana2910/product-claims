# product-claims

A Go backend that scans a product and identifies marketing or compliance claims from it.

tech stack: Go, PostgreSQL, net/http, google/uuid, joho/godotenv, lib/pq

prerequisites: Go 1.21+ and PostgreSQL 14+


setup

1. clone the repo
   git clone https://github.com/Sanjana2910/product-claims.git
   cd product-claims

2. install dependencies
   go mod tidy

3. create .env file in root
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=product_claims

4. set up the database
   psql -U postgres -c "CREATE DATABASE product_claims;"
   psql -U postgres -d product_claims -f db/migrations/1.sql
   psql -U postgres -d product_claims -f db/migrations/2.sql

5. run the server
   go run main.go


APIs

POST /claims/identify
body: { "productId": "cfe6aa75-5da8-44f5-b587-56857841ad9f" }
returns: { "workflowId": "wf-...", "status": "IN_PROGRESS" }

GET /claims/status/{workflowId}
returns the full product with identified claims once completed


tests

go test ./tests/ -v


postman collection is inside the postman/ folder
