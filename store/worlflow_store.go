package store
// talks directly with pstgres
import (
	"database/sql"
	"encoding/json"// convert json to go struct
	"fmt"
	"product-claims/models"
)
// talks with postgres 
//fetches a product from the database by its id
func GetProductByID(db *sql.DB, productID string) (*models.Product, error) {// returns the product if found and if not err
	var rawJSON string

	err := db.QueryRow(// gets one row
		"SELECT raw_json FROM products WHERE id = $1",
		productID,
	).Scan(&rawJSON) //scan puts the result into rawJSON variable

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("product not found: %s", productID)
	}
	if err != nil {//if something else went wrong with the query
		return nil, fmt.Errorf("error fetching product: %v", err)
	}
	// convert the raw JSON string into a Product struct
    var product models.Product
	err = json.Unmarshal([]byte(rawJSON), &product) // unmarshall needs bytes
	if err != nil {
		return nil, fmt.Errorf("error parsing product json: %v", err)
	}

	return &product, nil
}

// CreateWorkflow saves a new workflow in the database
func CreateWorkflow(db *sql.DB, workflowID, productID, status string) error {
	_, err := db.Exec(
		`INSERT INTO workflows (id, product_id, status) 
		 VALUES ($1, $2, $3)`,
		workflowID, productID, status,
	)
	return err
}

//updates the status of a workflow
func UpdateWorkflowStatus(db *sql.DB, workflowID, status, errorMsg string) error {
	_, err := db.Exec(
		`UPDATE workflows 
		 SET status = $1, error_msg = $2, updated_at = NOW() //postgres fills 
		 WHERE id = $3`,
		status, errorMsg, workflowID,
	)// updates the status 
	return err
}

//saves all found claims into the database
func SaveClaims(db *sql.DB, workflowID, productID string, claims []models.Claim) error {
	for _, claim := range claims {
		_, err := db.Exec(
			`INSERT INTO claims (id, workflow_id, product_id, claim_type, claim_value, status)
			 VALUES ($1, $2, $3, $4, $5, $6)`,
			claim.ID, workflowID, productID,
			claim.ClaimType, claim.ClaimValue, claim.Status,
		)
		if err != nil {
			return fmt.Errorf("error saving claim: %v", err)
		}
	}
	return nil
}

// GetWorkflowByID fetches a workflow and its claims from the database
func GetWorkflowByID(db *sql.DB, workflowID string) (*models.Workflow, error) {
	var workflow models.Workflow
	var errMsg sql.NullString

	err := db.QueryRow(
		`SELECT id, product_id, status, error_msg, created_at, updated_at
		 FROM workflows WHERE id = $1`,
		workflowID,
	).Scan(
		&workflow.ID,
		&workflow.ProductID,
		&workflow.Status,
		&errMsg,
		&workflow.CreatedAt,
		&workflow.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("workflow not found: %s", workflowID)
	}
	if err != nil {
		return nil, fmt.Errorf("error fetching workflow: %v", err)
	}

	if errMsg.Valid {
		workflow.ErrorMsg = errMsg.String
	}

	// if completed, fetch the claims too
	if workflow.Status == models.StatusCompleted {
		claims, err := GetClaimsByWorkflowID(db, workflowID)
		if err != nil {
			return nil, err
		}

		product, err := GetProductByID(db, workflow.ProductID)
		if err != nil {
			return nil, err
		}

		product.Claims = claims
		workflow.Product = product
	}

	return &workflow, nil
}

// GetClaimsByWorkflowID fetches all claims for a workflow
func GetClaimsByWorkflowID(db *sql.DB, workflowID string) ([]models.Claim, error) {
	rows, err := db.Query(
		`SELECT id, claim_type, claim_value, status
		 FROM claims WHERE workflow_id = $1`,
		workflowID,
	)
	if err != nil {
		return nil, fmt.Errorf("error fetching claims: %v", err)
	}
	defer rows.Close()

	var claims []models.Claim
	for rows.Next() {
		var claim models.Claim
		err := rows.Scan(
			&claim.ID,
			&claim.ClaimType,
			&claim.ClaimValue,
			&claim.Status,
		)
		if err != nil {
			return nil, err
		}
		claims = append(claims, claim)
	}

	return claims, nil
}
