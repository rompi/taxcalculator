package constant

var (
	FOOD_BEVERAGES = "Food & Beverages"
	TOBACCO        = "Tobacco"
	ENTERTAINMENT  = "Entertainment"

	// Map enum from tax code to string type
	TAXCODE_TYPE = map[int]string{
		1: FOOD_BEVERAGES,
		2: TOBACCO,
		3: ENTERTAINMENT,
	}

	// Object refundable
	IS_REFUNDABLE = map[int]bool{
		1: true,
		2: false,
		3: false,
	}
)
