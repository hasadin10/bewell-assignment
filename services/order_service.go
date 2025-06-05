package services

import (
	"bewell-app/entities"
	"bewell-app/utils"

	// "fmt"
	"strings"
)

func CleanedOrder(inputOrders []entities.InputOrder) []entities.CleanedOrder {
	var cleanedOrders []entities.CleanedOrder
	for _, order := range inputOrders {

		// แยก productId หลายตัว
		productIdStrs := utils.SplitProductIds(order.PlatformProductId)

		// วนลูปแต่ละ productId
		var totalQty int
		var productInfos []struct {
			ProductId  string
			MaterialId string
			ModelId    string
			Qty        int
		}

		// 1. วนลูปทุก inputOrder
		for _, p := range productIdStrs {
			productId, qty := utils.ExtractQtyFromProductId(p)              //แยก productId และ qty
			materialId, modelId := utils.ExtractMaterialAndModel(productId) //แยก materialId และ modelId
			totalQty += qty                                                 //นับจำนวน qty
			productInfos = append(productInfos, struct {
				ProductId  string
				MaterialId string
				ModelId    string
				Qty        int
			}{productId, materialId, modelId, qty}) //เพิ่ม productId, materialId, modelId, qty เข้าไปใน productInfos
		}

		// สร้าง response สำหรับแต่ละ product
		var uPrice, tPrice float64
		var qty int
		materialCount := map[string]int{} //สร้าง map เพื่อนับจำนวน materialId
		materialQty := map[string]int{}   //สร้าง map เพื่อนับจำนวน materialId
		for i, info := range productInfos {
			materialCount[info.MaterialId]++
			materialQty[info.MaterialId] += info.Qty
			if totalQty == 1 {
				// ถ้ามี productId เดียว ใช้ qty จาก input
				qty = inputOrders[0].Qty
				uPrice = inputOrders[0].UnitPrice
				tPrice = inputOrders[0].TotalPrice
			} else {
				// ถ้าแยกได้หลาย productId ให้ใช้ qty ที่แยกได้
				qty = info.Qty
				uPrice = utils.Round2(inputOrders[0].TotalPrice / float64(totalQty))
				tPrice = utils.Round2(uPrice * float64(qty))
			}
			cleanedOrders = append(cleanedOrders, entities.CleanedOrder{
				No:         i + 1,
				ProductId:  info.ProductId,
				MaterialId: info.MaterialId,
				ModelId:    info.ModelId,
				Qty:        qty,
				UnitPrice:  uPrice,
				TotalPrice: tPrice,
			})
		}

		// ของแถม ใช้ qty รวม
		var realQty int
		freeNo := len(productInfos) + 1 // เริ่มต่อจาก product หลัก
		if len(productInfos) == 1 {
			// ถ้า product เดียว แต่ qty > 1 (เช่น *3)
			if inputOrders[0].Qty > 1 {
				realQty = productInfos[0].Qty * inputOrders[0].Qty
			} else {
				realQty = productInfos[0].Qty
			}
			// กรณีเดียว
			cleanedOrders = append(cleanedOrders, entities.CleanedOrder{
				No:         freeNo,
				ProductId:  "WIPING-CLOTH",
				Qty:        realQty,
				UnitPrice:  0,
				TotalPrice: 0,
			})
			freeNo++
			for mat := range materialCount {
				var freeProduct string
				if strings.Contains(mat, "CLEAR") {
					freeProduct = "CLEAR-CLEANNER"
				} else if strings.Contains(mat, "MATTE") {
					freeProduct = "MATTE-CLEANNER"
				} else if strings.Contains(mat, "PRIVACY") {
					freeProduct = "PRIVACY-CLEANER"
				} else {
					continue
				}
				cleanedOrders = append(cleanedOrders, entities.CleanedOrder{
					No:         freeNo,
					ProductId:  freeProduct,
					Qty:        realQty,
					UnitPrice:  0,
					TotalPrice: 0,
				})
				freeNo++
			}
		} else {

			// กรณีหลาย product
			// รวม qty ของทุก product
			totalProductQty := 0
			for _, info := range productInfos {
				totalProductQty += info.Qty
			}
			cleanedOrders = append(cleanedOrders, entities.CleanedOrder{
				No:         freeNo,
				ProductId:  "WIPING-CLOTH",
				Qty:        totalProductQty,
				UnitPrice:  0,
				TotalPrice: 0,
			})
			freeNo++
			for mat := range materialCount {
				var freeProduct string
				if strings.Contains(mat, "CLEAR") {
					freeProduct = "CLEAR-CLEANNER"
				} else if strings.Contains(mat, "MATTE") {
					freeProduct = "MATTE-CLEANNER"
				} else if strings.Contains(mat, "PRIVACY") {
					freeProduct = "PRIVACY-CLEANER"
				} else {
					continue
				}

				cleanedOrders = append(cleanedOrders, entities.CleanedOrder{
					No:         freeNo,
					ProductId:  freeProduct,
					Qty:        materialQty[mat], // จำนวน product ที่เป็น materialId นั้น
					UnitPrice:  0,
					TotalPrice: 0,
				})
				freeNo++
			}
		}
	}
	return cleanedOrders
}
