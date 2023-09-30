package rating

import (
	"mood-tracker/src/rating/repository"
)

func AddRating(value int) (int64, error) {
	return repository.InsertRating(value)
}

func GetAllRatings() ([]repository.Rating, error) {
	return repository.GetAllRatings()
}

func UpdateRaring(id int, value int) error {
	return repository.UpdateRating(id, value)
}
