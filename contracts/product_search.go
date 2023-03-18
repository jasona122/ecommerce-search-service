package contracts

type ESProductSearchResult struct {
	ID     string                      `json:"_id"`
	Score  float64                     `json:"_score"`
	Source ESProductSearchResultSource `json:"_source"`
}

type ESProductSearchResultSource struct {
	Name          string          `json:"name"`
	Category      ProductCategory `json:"category"`
	Description   string          `json:"description"`
	ImageURL      string          `json:"image_url"`
	Price         float32         `json:"price"`
	Quantity      int             `json:"quantity"`
	ServiceAreaID string          `json:"service_area_id"`
	ShopName      string          `json:"shop_name"`
}

type ProductSearchResult struct {
	Name          string          `json:"name"`
	Category      ProductCategory `json:"category"`
	Description   string          `json:"description"`
	ImageURL      string          `json:"image_url"`
	Price         float32         `json:"price"`
	Quantity      int             `json:"quantity"`
	ServiceAreaID string          `json:"service_area_id"`
	ShopName      string          `json:"shop_name"`
}

type ProductSearchResponse struct {
	Results []ProductSearchResult `json:"results"`
}

func (response ProductSearchResponse) DataMarker() {

}
