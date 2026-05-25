# Product Claim Identification

A Go backend service that reads product data, detects marketing or compliance-related claims, classifies them into predefined categories, and returns the enriched product JSON with identified claims.

---

## What it does

When you give it a product ID, it:
1. Fetches the product from the database
2. Scans the product text fields for claims
3. Classifies each claim into a category
4. Saves the claims and returns them with the product



## Tech Stack

- Go — backend language
- PostgreSQL — database
- net/http — HTTP routing
- google/uuid — unique ID generation
- joho/godotenv — environment variable loading
- lib/pq — PostgreSQL driver


## Project Structure

product-claims/
├── main.go                      
├── .env                         
├── go.mod                       
│
├── db/
│   ├── db.go                    
│   └── migrations
|
│       ├── 1.sql         
│       └── 2.sql         
│
├── models/
│   ├── product.go               
│   ├── claim.go                 
│   └── workflow.go              
│
├── services/
│   ├── detector.go              
│   └── classifier.go            
│
├── store/
│   └── workflow_store.go        
│
├── handlers/
│   ├── claims.go                
│   └── status.go                
│
└── tests/
    ├── db_test.go               
    ├── classifier_test.go       
    ├── detector_test.go         
    ├── handler_test.go          
    └── store_test.go            


## Setup Instructions

1. Clone the repository

git clone https://github.com/yourusername/product-claims.git
cd product-claims

2. Install dependencies

go mod tidy

3. Create the .env file

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=product_claims

4. Set up the database

psql -U postgres -c "CREATE DATABASE product_claims;"
psql -U postgres -d product_claims -f db/migrations/001_init.sql
psql -U postgres -d product_claims -f db/migrations/002_seed.sql

5. Run the server

go run main.go

You should see:
postgres connected!
server running on http://localhost:8080

---

## API Reference

1. Trigger Claim Identification

POST /claims/identify

Request:
{
  "productId": "cfe6aa75-5da8-44f5-b587-56857841ad9f"
}

Response:
{
  "workflowId": "wf-1e7f6a3d-c855-4ce0-9a7c-45e3d1927356",
  "status": "IN_PROGRESS"
}

Error responses:
400 → productId is missing or empty
404 → product not found
500 → internal server error

---

2. Check Claim Status

GET /claims/status/{workflowId}

Example:
GET /claims/status/wf-1e7f6a3d-c855-4ce0-9a7c-45e3d1927356

Response when IN_PROGRESS:
{
  "workflowId": "wf-1e7f6a3d-...",
  "status": "IN_PROGRESS",
  "createdAt": "2026-01-13T15:29:26Z",
  "updatedAt": "2026-01-13T15:29:26Z"
}

Response when COMPLETED:
{
  "workflowId": "wf-1e7f6a3d-...",
  "status": "COMPLETED",
  "product": {
    "id": "cfe6aa75-...",
    "title": "Maggi Nutri-licious Masala Veg Atta Noodles",
    "brand": "Maggi",
    "claims": [
      {
        "id": "uuid",
        "claimType": "Nutritional Claims",
        "claimValue": "Source of Fibre & Iron",
        "status": "IDENTIFIED"
      },
      {
        "id": "uuid",
        "claimType": "Safety Claims",
        "claimValue": "Safe for children",
        "status": "IDENTIFIED"
      },
      {
        "id": "uuid",
        "claimType": "Manufacturing / Origin",
        "claimValue": "Made in India",
        "status": "IDENTIFIED"
      },
      {
        "id": "uuid",
        "claimType": "Certification Claims",
        "claimValue": "FSSAI certified",
        "status": "IDENTIFIED"
      }
    ]
  }
}

Response when FAILED:
{
  "workflowId": "wf-1e7f6a3d-...",
  "status": "FAILED",
  "errorMsg": "product not found"
}

Error responses:
400 → workflowId missing from URL
404 → workflow not found

---

## Claim Categories

Performance / Efficacy        → Long-lasting, More effective, High strength
Health & Wellness             → Supports immunity, Promotes digestion
Nutritional Claims            → High protein, Low sugar, Source of Fibre & Iron
Comparative Claims            → 30% better than ordinary, Uses less energy
Environmental / Sustainability → Eco-friendly, Recyclable, Sustainable
Safety Claims                 → Safe for children, Non-toxic, Dermatologically tested
Quality Claims                → Premium quality, High-grade materials
Manufacturing / Origin        → Made in India, Handcrafted, Artisan made
Certification Claims          → FSSAI certified, BIS approved, ISO compliant
Awards & Endorsements         → Award-winning, Recognized by XYZ
Price / Value Claims          → Best value, Lowest price, Big savings
Authenticity Claims           → 100% original, Authentic, Traditional recipe
Medical / Therapeutic         → Cures diabetes, Treats arthritis

---

## Running Tests

go test ./tests/ -v

Expected output:
--- PASS: TestNutritionalClaim
--- PASS: TestSafetyClaim
--- PASS: TestCertificationClaim
--- PASS: TestManufacturingClaim
--- PASS: TestNoMatchClaim
--- PASS: TestDBConnection
--- PASS: TestMaggiProductExists
--- PASS: TestDetectBasicClaims
--- PASS: TestDetectNoDuplicates
--- PASS: TestDetectEmptyProduct
--- PASS: TestGetMaggiProduct
--- PASS: TestGetInvalidProduct
--- PASS: TestCreateWorkflow
--- PASS: TestSaveAndFetchClaims
PASS

---

## Database Schema

products table:
id, mcr_id, title, brand, category_name, raw_json, created_at, updated_at

workflows table:
id, product_id, status, error_msg, created_at, updated_at

claims table:
id, workflow_id, product_id, claim_type, claim_value, status, created_at

---

## Sample curl Commands

# trigger claim identification
curl -X POST http://localhost:8080/claims/identify \
  -H "Content-Type: application/json" \
  -d '{"productId": "cfe6aa75-5da8-44f5-b587-56857841ad9f"}'

# check status
curl http://localhost:8080/claims/status/wf-1e7f6a3d-c855-4ce0-9a7c-45e3d1927356

---

## Assumptions

- Claim detection is based on keyword matching against predefined patterns
- A product must exist in the database before triggering claim identification
- Claim detection runs asynchronously using goroutines
- Duplicate claims from the same product are filtered out automatically
- FSSAI number in complianceInfo is treated as a Certification Claim
- Country of origin is treated as a Manufacturing / Origin claim
- Workflow status moves from IN_PROGRESS → COMPLETED or FAILED