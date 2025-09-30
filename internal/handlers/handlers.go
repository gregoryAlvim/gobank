package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/gregoryAlvim/gobank/internal/services"
)

type AccountHandler struct {
	service *services.AccountService
}

func NewAccountHandler(service *services.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	accountType := r.URL.Query().Get("type")
	if accountType == "" {
		http.Error(w, "Account type is required", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	err = h.service.CreateAccount(accountType, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Account created successfully"})
}

func (h *AccountHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}
	accountType := r.URL.Query().Get("type")
	if accountType == "" {
		http.Error(w, "Account type is required", http.StatusBadRequest)
		return
	}

	balance, err := h.service.GetBalance(id, accountType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"balance": balance})
}

type AmountRequest struct {
	Amount float64 `json:"amount"`
}

func (h *AccountHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}
	accountType := r.URL.Query().Get("type")
	if accountType == "" {
		http.Error(w, "Account type is required", http.StatusBadRequest)
		return
	}

	var req AmountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.Deposit(id, req.Amount, accountType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Deposit successful"})
}

func (h *AccountHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}
	accountType := r.URL.Query().Get("type")
	if accountType == "" {
		http.Error(w, "Account type is required", http.StatusBadRequest)
		return
	}

	var req AmountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.Withdraw(id, req.Amount, accountType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Withdrawal successful"})
}

type TransferRequest struct {
	FromID   int     `json:"from_id"`
	ToID     int     `json:"to_id"`
	FromType string  `json:"from_type"`
	ToType   string  `json:"to_type"`
	Amount   float64 `json:"amount"`
}

func (h *AccountHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	var req TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.Transfer(req.FromID, req.ToID, req.Amount, req.FromType, req.ToType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Transfer successful"})
}

func (h *AccountHandler) CloseAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}
	accountType := r.URL.Query().Get("type")
	if accountType == "" {
		http.Error(w, "Account type is required", http.StatusBadRequest)
		return
	}

	if err := h.service.CloseAccount(id, accountType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Account closed successfully"})
}
