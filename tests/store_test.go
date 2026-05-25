package tests

import (
	"testing"
	"product-claims/db"
	"product-claims/models"
	"product-claims/store"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func setupStore() {
	godotenv.Load("../.env")
	db.Connect()
}

func TestGetMaggiProduct(t *testing.T) {
	setupStore()

	maggiID := "cfe6aa75-5da8-44f5-b587-56857841ad9f"
	product, err := store.GetProductByID(db.Conn, maggiID)
	if err != nil {
		t.Fatalf("expected product, got error: %v", err)
	}

	if product.Title == "" {
		t.Fatal("expected title, got empty")
	}

	t.Logf("got product: %s", product.Title)
}

func TestGetInvalidProduct(t *testing.T) {
	setupStore()

	_, err := store.GetProductByID(db.Conn, "invalid-id")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	t.Log("invalid id returns error!")
}

func TestCreateWorkflow(t *testing.T) {
	setupStore()

	wfID      := "wf-" + uuid.New().String()
	productID := "cfe6aa75-5da8-44f5-b587-56857841ad9f"

	err := store.CreateWorkflow(db.Conn, wfID, productID, models.StatusInProgress)
	if err != nil {
		t.Fatalf("could not create workflow: %v", err)
	}

	wf, err := store.GetWorkflowByID(db.Conn, wfID)
	if err != nil {
		t.Fatalf("could not fetch workflow: %v", err)
	}

	if wf.Status != models.StatusInProgress {
		t.Fatalf("expected IN_PROGRESS, got %s", wf.Status)
	}

	t.Logf("workflow created and fetched: %s", wf.ID)
}

func TestSaveAndFetchClaims(t *testing.T) {
	setupStore()

	wfID      := "wf-" + uuid.New().String()
	productID := "cfe6aa75-5da8-44f5-b587-56857841ad9f"

	store.CreateWorkflow(db.Conn, wfID, productID, models.StatusInProgress)

	claims := []models.Claim{
		{
			ID:         uuid.New().String(),
			ClaimType:  "Nutritional Claims",
			ClaimValue: "Source of Fibre & Iron",
			Status:     "IDENTIFIED",
		},
		{
			ID:         uuid.New().String(),
			ClaimType:  "Safety Claims",
			ClaimValue: "Safe for children",
			Status:     "IDENTIFIED",
		},
	}

	err := store.SaveClaims(db.Conn, wfID, productID, claims)
	if err != nil {
		t.Fatalf("could not save claims: %v", err)
	}

	t.Logf("saved %d claims!", len(claims))
}
