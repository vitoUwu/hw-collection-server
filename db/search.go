package db

import (
	"strings"
)

type Search struct {
	ID         string `json:"id"`
	ModelName  string `json:"model_name"`
	Code       string `json:"code"`
	SeriesName string `json:"series_name"`
	SeriesCode string `json:"series_code"`
	Notes      string `json:"notes"`
	ColCode    string `json:"col_code"`
}

type SearchResult struct {
	ID         string `json:"id"`
	ModelName  string `json:"model_name"`
	Code       string `json:"code"`
	SeriesName string `json:"series_name"`
	SeriesCode string `json:"series_code"`
	Notes      string `json:"notes"`
	ColCode    string `json:"col_code"`
	Image      string `json:"image"`
}

const sqlQuery = ` 
	SELECT searches.id, searches.model_name, searches.code, searches.series_name, searches.series_code, searches.notes, searches.col_code, cars.image
	FROM searches 
	INNER JOIN cars ON searches.id = cars.id
	WHERE searches.id % ? 
	OR searches.code % ? 
	OR searches.model_name % ? 
	OR searches.series_name % ? 
	OR searches.series_code % ? 
	OR searches.notes % ? 
	OR searches.col_code % ?
	ORDER BY GREATEST(
			similarity(searches.id, ?),
			similarity(searches.code, ?),
			similarity(searches.model_name, ?),
			similarity(searches.series_name, ?),
			similarity(searches.series_code, ?),
			similarity(searches.notes, ?),
			similarity(searches.col_code, ?)
	) DESC
`

func (db *Database) Search(query string) ([]SearchResult, error) {
	params := []interface{}{}
	countParams := strings.Count(sqlQuery, "?")

	for i := 0; i < countParams; i++ {
		params = append(params, query)
	}

	var results []SearchResult
	err := db.Gorm.Raw(sqlQuery, params...).Scan(&results).Error

	return results, err
}
