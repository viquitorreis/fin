package main

type Response struct {
	Indexes             []Index  `json:"indexes"`
	Stocks              []Stock  `json:"stocks"`
	AvailableSectors    []string `json:"availableSectors"`
	AvailableStockTypes []string `json:"availableStockTypes"`
	CurrentPage         int      `json:"currentPage"`
	TotalPages          int      `json:"totalPages"`
	ItemsPerPage        int      `json:"itemsPerPage"`
	TotalItems          int      `json:"totalItems"`
}

type Index struct {
	Stock string `json:"stock"`
	Name  string `json:"name"`
}

type Stock struct {
	Stock     string  `json:"stock"`
	Name      string  `json:"name"`
	Close     float64 `json:"close"`
	Change    float64 `json:"change"`
	Volume    int     `json:"volume"`
	MarketCap float64 `json:"marketCap"`
	Logo      string  `json:"logo"`
	Sector    string  `json:"sector"`
	Type      string  `json:"type"`
}
