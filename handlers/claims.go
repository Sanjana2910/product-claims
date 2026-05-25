package handlers
// handles the incoming api requests
import (
	"encoding/json"
	"log"
	"net/http"
	"product-claims/db"
	"product-claims/models"
	"product-claims/services"
	"product-claims/store"

	"github.com/google/uuid"
)
// req and response structures

//what the user sends in the request body
type TriggerRequest struct {
	ProductID string `json:"productId"`
}

//whats sent back
type TriggerResponse struct {
	WorkflowID string `json:"workflowId"`
	Status     string `json:"status"`
}

//API function that handles requests and responses (POST /claims/identify)
func IdentifyClaims(w http.ResponseWriter, r *http.Request) { // r,w from net/http
	// only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// read the request body 
	var req TriggerRequest
	err := json.NewDecoder(r.Body).Decode(&req)// creates json reader and fills value inside req
	if err != nil || req.ProductID == "" {
		http.Error(w, "invalid request, productId is required", http.StatusBadRequest)
		return
	}

	// check if product exists
	_, err = store.GetProductByID(db.Conn, req.ProductID)
	if err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	// create a new workflow id
	workflowID := "wf-" + uuid.New().String()

	// save the workflow as IN_PROGRESS
	err = store.CreateWorkflow(db.Conn, workflowID, req.ProductID, models.StatusInProgress)
	if err != nil {
		http.Error(w, "could not create workflow", http.StatusInternalServerError)
		return
	}

	// run the actual work in the background
	// goroutine means it runs without making the user wait
	go runClaimDetection(workflowID, req.ProductID)

	// immediately return the workflow id and status
	w.Header().Set("Content-Type", "application/json")//sending json data
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(TriggerResponse{ //convert go to json
		WorkflowID: workflowID,
		Status: models.StatusInProgress,
	})
}

// runClaimDetection does the actual work in the background
func runClaimDetection(workflowID, productID string) {
	// fetch the product
	product, err := store.GetProductByID(db.Conn, productID)
	if err != nil {
		log.Printf("error fetching product: %v", err)
		store.UpdateWorkflowStatus(db.Conn, workflowID, models.StatusFailed, err.Error())
		return
	}

	// detect and classify claims
	claims := services.Detect(product)

	// give each claim a unique id
	for i := range claims {
		claims[i].ID = uuid.New().String()
	}

	// save claims to database
	err = store.SaveClaims(db.Conn, workflowID, productID, claims)
	if err != nil {
		log.Printf("error saving claims: %v", err)
		store.UpdateWorkflowStatus(db.Conn, workflowID, models.StatusFailed, err.Error())
		return
	}

	// mark workflow as completed
	store.UpdateWorkflowStatus(db.Conn, workflowID, models.StatusCompleted, "")
	log.Printf("workflow %s completed with %d claims", workflowID, len(claims))
}