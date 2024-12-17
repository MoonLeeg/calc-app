package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MoonLeeg/calc-app/internal/calculator"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Invalid method"}`, http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request"}`, http.StatusBadRequest)
		return
	}

	if req.Expression == "" {
		http.Error(w, `{"error": "Expression is not valid"}`, http.StatusUnprocessableEntity)
		return
	}

	result, err := calculator.Calc(req.Expression)
	if err != nil {
		var status int
		if err.Error() == "division by zero" {
			status = http.StatusUnprocessableEntity
		} else {
			status = http.StatusUnprocessableEntity
		}
		http.Error(w, `{"error": "`+err.Error()+`"}`, status)
		return
	}

	resp := Response{Result: resultToString(result)}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func resultToString(result float64) string {
	return strconv.FormatFloat(result, 'f', -1, 64)
}
