package utils

import (
	"math"
	// "regexp"
    "strings"
    "strconv"
    // "fmt"
)

func Round2(f float64) float64 {
    return math.Round(f*100) / 100
}



func ExtractProductId(raw string) string {
    idx := strings.Index(raw, "FG")
    if idx == -1 {
        return "" // ไม่เจอ FG เลย
    }
    return raw[idx:]
}

func ExtractQtyFromProductId(raw string) (productId string, qtyFromId int) {
    productId = ExtractProductId(raw)
    qtyFromId = 1
    if strings.Contains(raw, "*") {
        parts := strings.Split(raw, "*")
        productId = ExtractProductId(parts[0])
        if n, err := strconv.Atoi(parts[1]); err == nil {
            qtyFromId = n
        }
    }

    return productId, qtyFromId
}

// แยก productId หลายตัวจาก string ที่คั่นด้วย /
func SplitProductIds(raw string) []string {
    parts := strings.Split(raw, "/")
    var result []string
    for _, p := range parts {
        p = strings.TrimSpace(p)
        if p != "" {
            result = append(result, p)
        }
    }
    return result
}


func ExtractMaterialAndModel(productId string) (materialId, modelId string) {
    parts := strings.Split(productId, "-")
    if len(parts) >= 3 {
        materialId = parts[0] + "-" + parts[1]
        modelId = strings.Join(parts[2:], "-")
    }
    return
}