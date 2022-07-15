package types

// OpenseaStats
//{
//  "stats": {
//    "one_day_volume": 63.27036099900001,
//    "one_day_change": -0.5299071128428349,
//    "one_day_sales": 202,
//    "one_day_average_price": 0.31321960890594064,
//    "seven_day_volume": 634.0186629669897,
//    "seven_day_change": 0,
//    "seven_day_sales": 2107,
//    "seven_day_average_price": 0.3009106136530563,
//    "thirty_day_volume": 634.0186629669897,
//    "thirty_day_change": 0,
//    "thirty_day_sales": 2107,
//    "thirty_day_average_price": 0.3009106136530563,
//    "total_volume": 634.0186629669897,
//    "total_sales": 2107,
//    "total_supply": 5500,
//    "count": 5500,
//    "num_owners": 1804,
//    "average_price": 0.3009106136530563,
//    "num_reports": 1,
//    "market_cap": 1655.00837509181,
//    "floor_price": 0.293
//  }
//}
type OpenseaStats struct {
	Stats struct {
		OneDayVolume          float64 `json:"one_day_volume"`
		OneDayChange          float64 `json:"one_day_change"`
		OneDaySales           float64 `json:"one_day_sales"`
		OneDayAveragePrice    float64 `json:"one_day_average_price"`
		SevenDayVolume        float64 `json:"seven_day_volume"`
		SevenDayChange        float64 `json:"seven_day_change"`
		SevenDaySales         float64 `json:"seven_day_sales"`
		SevenDayAveragePrice  float64 `json:"seven_day_average_price"`
		ThirtyDayVolume       float64 `json:"thirty_day_volume"`
		ThirtyDayChange       float64 `json:"thirty_day_change"`
		ThirtyDaySales        float64 `json:"thirty_day_sales"`
		ThirtyDayAveragePrice float64 `json:"thirty_day_average_price"`
		TotalVolume           float64 `json:"total_volume"`
		TotalSales            float64 `json:"total_sales"`
		TotalSupply           float64 `json:"total_supply"`
		Count                 float64 `json:"count"`
		NumOwners             float64 `json:"num_owners"`
		AveragePrice          float64 `json:"average_price"`
		NumReports            float64 `json:"num_reports"`
		MarketCap             float64 `json:"market_cap"`
		FloorPrice            float64 `json:"floor_price"`
	} `json:"stats"`
}
