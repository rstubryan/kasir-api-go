package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/service"
	"net/http"

	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(service service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}

// RegisterRoutes registers transaction routes
func (h *TransactionHandler) RegisterRoutes(router *mux.Router) {
	// Checkout endpoint
	router.HandleFunc("/api/checkout", h.HandleCheckout).Methods(http.MethodPost)

	// Report endpoints
	router.HandleFunc("/api/report/hari-ini", h.GetDailyReport).Methods(http.MethodGet)
	router.HandleFunc("/api/report", h.GetReportByDateRange).Methods(http.MethodGet)
}

// HandleCheckout handles POST /api/checkout
func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	// Validate items
	if len(req.Items) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Items cannot be empty",
		})
		return
	}

	transaction, err := h.service.CreateTransaction(req.Items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

// GetDailyReport handles GET /api/report/hari-ini
func (h *TransactionHandler) GetDailyReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	report, err := h.service.GetDailyReport()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to get daily report",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(report)
}

// GetReportByDateRange handles GET /api/report?start_date=...&end_date=...
func (h *TransactionHandler) GetReportByDateRange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get query parameters
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	// If no date range provided, return bad request
	if startDate == "" || endDate == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "start_date and end_date query parameters are required",
		})
		return
	}

	report, err := h.service.GetReportByDateRange(startDate, endDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to get report",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(report)
}
