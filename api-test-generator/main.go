package main

import (
    "encoding/json"
    "net/http"
)

type FieldRule struct {
    Required bool   `json:"required"`
    MinLen   int    `json:"min_len"`
    MaxLen   int    `json:"max_len"`
    Type     string `json:"type"`
}

type Field struct {
    Name  string    `json:"name"`
    Rules FieldRule `json:"rules"`
}

type TestCase struct {
    Description string      `json:"description"`
    Value       interface{} `json:"value"`
}

func validateAndGenerateTestCases(fields []Field) []TestCase {
    var testCases []TestCase

    for _, field := range fields {
        // 正常情况
        testCases = append(testCases, TestCase{
            Description: field.Name + " normal case",
            Value:       generateNormalValue(field),
        })

        // 边界情况和错误情况
        if field.Rules.Required {
            testCases = append(testCases, TestCase{
                Description: field.Name + " missing (required)",
                Value:       nil,
            })
        }

        if field.Rules.MinLen != 0 {
            testCases = append(testCases, TestCase{
                Description: field.Name + " too short",
                Value:       generateTooShortValue(field),
            })
        }

        if field.Rules.MaxLen != 0 {
            testCases = append(testCases, TestCase{
                Description: field.Name + " too long",
                Value:       generateTooLongValue(field),
            })
        }

        if field.Rules.Type == "int" {
            testCases = append(testCases, TestCase{
                Description: field.Name + " invalid type (not int)",
                Value:       "string",
            })
        }
    }

    return testCases
}

func generateNormalValue(field Field) interface{} {
    switch field.Rules.Type {
    case "string":
        return "normal"
    case "int":
        return 123
    }
    return nil
}

func generateTooShortValue(field Field) interface{} {
    if field.Rules.Type == "string" {
        return ""
    }
    return nil
}

func generateTooLongValue(field Field) interface{} {
    switch field.Rules.Type {
    case "string":
        // 生成一个长度为 max_len + 1 的字符串
        return generateStringOfLength(field.Rules.MaxLen + 1)
    case "int":
        // 生成一个超出长度限制的整数，假设整数为固定长度的数字（可以根据实际需要调整）
        return generateIntExceedingMaxLen(field.Rules.MaxLen)
    }
    return nil
}

// 辅助函数：生成长度为指定长度的字符串
func generateStringOfLength(length int) string {
    if length <= 0 {
        return ""
    }
    // 生成长度为 length 的字符串，例如使用 'a'
    return string(make([]byte, length, 'a')) // 使用 make 创建一个指定长度的字节切片
}

// 辅助函数：生成超出指定长度的整数
func generateIntExceedingMaxLen(maxLen int) int {
    // 最大长度为 maxLen，生成一个超出这个长度的整数（例如 10^maxLen）
    return int(pow10(maxLen))
}

// 计算 10 的某次幂
func pow10(exp int) int64 {
    if exp < 0 {
        return 0
    }
    result := int64(1)
    for i := 0; i < exp; i++ {
        result *= 10
    }
    return result
}

func handleGenerateTestCases(w http.ResponseWriter, r *http.Request) {
    var fields []Field

    err := json.NewDecoder(r.Body).Decode(&fields)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    testCases := validateAndGenerateTestCases(fields)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(testCases)
}


func main() {
    http.HandleFunc("/generate-test-cases", handleGenerateTestCases)
    http.ListenAndServe(":8080", nil)
}
