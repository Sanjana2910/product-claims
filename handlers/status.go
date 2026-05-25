package handlers

import (
	"encoding/json"
	"net/http"
	"product-claims/db"
	"product-claims/store"
	"strings"
)

//handles GET 
func GetClaimStatus(w http.ResponseWriter, r *http.Request) {
	// only allow GET method
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// get the workflowId from the URL
	// URL looks like: /claims/status/wf-123
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 || parts[3] == "" {
		http.Error(w, "workflowId is required", http.StatusBadRequest)
		return
	}
	workflowID := parts[3] //like: workflowID = "wf-go-123"

	// fetch the workflow from database
	workflow, err := store.GetWorkflowByID(db.Conn, workflowID)
	if err != nil {
		http.Error(w, "workflow not found", http.StatusNotFound)
		return
	}

	// send back the workflow as json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workflow)//structs to json
}