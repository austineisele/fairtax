package handler

import (
  "fair-tax/internal/domain"
)

type Reponse struct {
	Name     string     `json:"name"`
	LevyData []LevyData `json:"rate_levy_entries"`
}
