package models

type BasicBook struct {
    Title         string   `json:"title"`
    Authors       []string `json:"author"`
    Category      string   `json:"category"`
    Subcategories []string `json:"subcategory"`
    ShelfLocation string   `json:"shelf_location"`
    ISBN          string   `json:"isbn"`
    Read          bool     `json:"read"`
    PageCount     int      `json:"page_count"`
}
