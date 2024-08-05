package handlers

import (
    "encoding/json"
    "net/http"
    "api-test-generator/models"
    "api-test-generator/validation"
)

func GenerateTestCases(w http.ResponseWriter, r *http.Request) {
    var req models.APIRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := validation.ValidateStruct(req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // 生成测试用例的逻辑
    response := map[string]string{"message": "Test cases generated successfully"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
