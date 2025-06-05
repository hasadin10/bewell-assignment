package services

import (
	"bewell-app/entities"
	"reflect"
	"testing"
)

// func TestCleanedOrder(t *testing.T) {
// 	input := []entities.InputOrder{
// 		{
// 			No:                1,
// 			PlatformProductId: "FG0A-CLEAR-IPHONE16PROMAX",
// 			Qty:               2,
// 			UnitPrice:         50,
// 			TotalPrice:        100,
// 		},
// 	}

// 	expected := []entities.CleanedOrder{
// 		{No: 1, ProductId: "FG0A-CLEAR-IPHONE16PROMAX", MaterialId: "FG0A-CLEAR", ModelId: "IPHONE16PROMAX", Qty: 2, UnitPrice: 50, TotalPrice: 100},
// 		{No: 2, ProductId: "WIPING-CLOTH", Qty: 2, UnitPrice: 0, TotalPrice: 0},
// 		{No: 3, ProductId: "CLEAR-CLEANNER", Qty: 2, UnitPrice: 0, TotalPrice: 0},
// 	}

// 	result := CleanedOrder(input)

// 	if !reflect.DeepEqual(result, expected) {
// 		t.Errorf("CleanedOrder() = %v, want %v", result, expected)
// 	}
// }

func TestCleanedOrder(t *testing.T) {
	tests := []struct {
		name     string
		input    []entities.InputOrder
		expected []entities.CleanedOrder
	}{
		{
			name: "Case 1 : Only one product",
			input: []entities.InputOrder{
				{
					No:                1,
					PlatformProductId: "FG0A-CLEAR-IPHONE16PROMAX",
					Qty:               2,
					UnitPrice:         50,
					TotalPrice:        100,
				},
			},
			expected: []entities.CleanedOrder{
				{No: 1, ProductId: "FG0A-CLEAR-IPHONE16PROMAX", MaterialId: "FG0A-CLEAR", ModelId: "IPHONE16PROMAX", Qty: 2, UnitPrice: 50, TotalPrice: 100},
				{No: 2, ProductId: "WIPING-CLOTH", Qty: 2, UnitPrice: 0, TotalPrice: 0},
				{No: 3, ProductId: "CLEAR-CLEANNER", Qty: 2, UnitPrice: 0, TotalPrice: 0},
			},
		},
		{
			name: "Case 2: One product with wrong prefix",
			input: []entities.InputOrder{
				{
					No:                1,
					PlatformProductId: "x2-3&FG0A-CLEAR-IPHONE16PROMAX",
					Qty:               2,
					UnitPrice:         50,
					TotalPrice:        100,
				},
			},
			expected: []entities.CleanedOrder{
				{No: 1, ProductId: "FG0A-CLEAR-IPHONE16PROMAX", MaterialId: "FG0A-CLEAR", ModelId: "IPHONE16PROMAX", Qty: 2, UnitPrice: 50, TotalPrice: 100},
				{No: 2, ProductId: "WIPING-CLOTH", Qty: 2, UnitPrice: 0, TotalPrice: 0},
				{No: 3, ProductId: "CLEAR-CLEANNER", Qty: 2, UnitPrice: 0, TotalPrice: 0},
			},
		},
		{
			name:  "Case 3: One product with wrong prefix and has * symbol that indicates the quantity",
			input: []entities.InputOrder{
				{
					No:                1,
					PlatformProductId: "x2-3&FG0A-MATTE-IPHONE16PROMAX*3",
					Qty:               1,
					UnitPrice:         90,
					TotalPrice:        90,
				},
			},
			expected: []entities.CleanedOrder{
				{No: 1, ProductId: "FG0A-MATTE-IPHONE16PROMAX", MaterialId: "FG0A-MATTE", ModelId: "IPHONE16PROMAX", Qty: 3, UnitPrice: 30, TotalPrice: 90},
				{No: 2, ProductId: "WIPING-CLOTH", Qty: 3, UnitPrice: 0, TotalPrice: 0},
				{No: 3, ProductId: "MATTE-CLEANNER", Qty: 3, UnitPrice: 0, TotalPrice: 0},
			},
		},
		{
			name:  "Case 4: One bundle product with wrong prefix and split by / symbol into two product",
			input: []entities.InputOrder{
				{
					No:                1,
					PlatformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B",
					Qty:               1,
					UnitPrice:         80,
					TotalPrice:        80,
				},
			},
			expected: []entities.CleanedOrder{
				{No: 1, ProductId: "FG0A-CLEAR-OPPOA3", MaterialId: "FG0A-CLEAR", ModelId: "OPPOA3", Qty: 1, UnitPrice: 40, TotalPrice: 40},
				{No: 2, ProductId: "FG0A-CLEAR-OPPOA3-B" ,MaterialId: "FG0A-CLEAR", ModelId: "OPPOA3-B", Qty: 1, UnitPrice: 40, TotalPrice: 40},
				{No: 3, ProductId: "WIPING-CLOTH", Qty: 2, UnitPrice: 0, TotalPrice: 0},
				{No: 4, ProductId: "CLEAR-CLEANNER", Qty: 2, UnitPrice: 0, TotalPrice: 0},
			},
		},
		{
			name:  "Case 5: One bundle product with wrong prefix and split by / symbol into three product",
			input: []entities.InputOrder{
				{
					No:                1,
					PlatformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B/FG0A-MATTE-OPPOA3",
					Qty:               1,
					UnitPrice:         120,
					TotalPrice:        120,
				},
			},
			expected: []entities.CleanedOrder{
				{No: 1, ProductId: "FG0A-CLEAR-OPPOA3", MaterialId: "FG0A-CLEAR", ModelId: "OPPOA3", Qty: 1, UnitPrice: 40, TotalPrice: 40},
				{No: 2, ProductId: "FG0A-CLEAR-OPPOA3-B" ,MaterialId: "FG0A-CLEAR", ModelId: "OPPOA3-B", Qty: 1, UnitPrice: 40, TotalPrice: 40},
				{No: 3, ProductId: "FG0A-MATTE-OPPOA3" ,MaterialId: "FG0A-MATTE", ModelId: "OPPOA3", Qty: 1, UnitPrice: 40, TotalPrice: 40},
				{No: 4, ProductId: "WIPING-CLOTH", Qty: 3, UnitPrice: 0, TotalPrice: 0},
				{No: 5, ProductId: "CLEAR-CLEANNER", Qty: 2, UnitPrice: 0, TotalPrice: 0},
				{No: 6, ProductId: "MATTE-CLEANNER", Qty: 1, UnitPrice: 0, TotalPrice: 0},
			},
		},
		{
			name:  "Case 6: One bundle product with wrong prefix and have / symbol and * symbol",
			input: []entities.InputOrder{
				{
					No:                1,
					PlatformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3",
					Qty:               1,
					UnitPrice:         120,
					TotalPrice:        120,
				},
			},
			expected: []entities.CleanedOrder{
				{No: 1, ProductId: "FG0A-CLEAR-OPPOA3", MaterialId: "FG0A-CLEAR", ModelId: "OPPOA3", Qty: 2, UnitPrice: 40, TotalPrice: 80},
				{No: 2, ProductId: "FG0A-MATTE-OPPOA3" ,MaterialId: "FG0A-MATTE", ModelId: "OPPOA3", Qty: 1, UnitPrice: 40, TotalPrice: 40},
				{No: 3, ProductId: "WIPING-CLOTH", Qty: 3, UnitPrice: 0, TotalPrice: 0},
				{No: 4, ProductId: "CLEAR-CLEANNER", Qty: 2, UnitPrice: 0, TotalPrice: 0},
				{No: 5, ProductId: "MATTE-CLEANNER", Qty: 1, UnitPrice: 0, TotalPrice: 0},
			},
		},
		{
			name:  "Case 7: one product and one bundle product with wrong prefix and have / symbol and * symbol",
			input: []entities.InputOrder{
				{
					No:                1,
					PlatformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3*2",
					Qty:               1,
					UnitPrice:         160,
					TotalPrice:        160,
				},
				{
					No:                2,
					PlatformProductId: "FG0A-PRIVACY-IPHONE16PROMAX",
					Qty:               1,
					UnitPrice:         50,
					TotalPrice:        50,
				},
			},
			expected: []entities.CleanedOrder{
				{No: 1, ProductId: "FG0A-CLEAR-OPPOA3", MaterialId: "FG0A-CLEAR", ModelId: "OPPOA3", Qty: 2, UnitPrice: 40, TotalPrice: 60},
				{No: 2, ProductId: "FG0A-MATTE-OPPOA3", MaterialId: "FG0A-MATTE", ModelId: "OPPOA3", Qty: 1, UnitPrice: 40, TotalPrice: 80},
				{No: 3, ProductId: "FG0A-PRIVACY-IPHONE16PROMAX", MaterialId: "FG0A-PRIVACY", ModelId: "IPHONE16PROMAX", Qty: 1, UnitPrice: 50, TotalPrice: 50},
				{No: 4, ProductId: "WIPING-CLOTH", Qty: 5, UnitPrice: 0, TotalPrice: 0},
				{No: 5, ProductId: "CLEAR-CLEANNER", Qty: 2, UnitPrice: 0, TotalPrice: 0},
				{No: 6, ProductId: "MATTE-CLEANNER", Qty: 2, UnitPrice: 0, TotalPrice: 0},
				{No: 7, ProductId: "PRIVACY-CLEANNER", Qty: 1, UnitPrice: 0, TotalPrice: 0},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := CleanedOrder(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Test %s failed: got %v, want %v", tc.name, result, tc.expected)
			}
		})
	}
}
