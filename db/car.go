package db

import "time"

type CarCharacteristics struct {
	ModelName     string  `json:"model_name"`
	Image         string  `json:"image"`
	Code          *string `json:"code"`
	SeriesName    *string `json:"series_name"`
	SeriesCode    *string `json:"series_code"`
	Notes         *string `json:"notes"`
	WheelType     *string `json:"wheel_type"`
	Color         *string `json:"color"`
	Tampo         *string `json:"tampo"`
	BaseColor     *string `json:"base_color"`
	WindowColor   *string `json:"window_color"`
	InteriorColor *string `json:"interior_color"`
	Country       *string `json:"country"`
	ColCode       *string `json:"col_code"`
}

type Car struct {
	ID string `json:"id"`

	CarCharacteristics `gorm:"embedded"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (db *Database) CreateCar(car *CarCharacteristics) (*Car, error) {
	newCar := Car{
		ID:                 db.GenerateId(),
		CarCharacteristics: *car,
	}

	err := db.Gorm.Model(&Car{}).Create(&newCar).Error

	return &newCar, err
}

func (db *Database) CreateManyCars(cars []CarCharacteristics) error {
	var newCars []Car

	for _, car := range cars {
		newCars = append(newCars, Car{
			ID:                 db.GenerateId(),
			CarCharacteristics: car,
		})
	}

	return db.Gorm.Model(&Car{}).Create(&newCars).Error
}

func (db *Database) GetCarById(id string) (*Car, error) {
	var car Car

	err := db.Gorm.Model(&Car{}).Where("id = ?", id).First(&car).Error

	return &car, err
}

func (db *Database) GetCars(limit int, offset int) ([]Car, error) {
	var cars []Car

	err := db.Gorm.Model(&Car{}).Limit(limit).Offset(offset).Find(&cars).Error

	return cars, err
}

func (db *Database) CountCars() (int64, error) {
	var count int64

	err := db.Gorm.Model(&Car{}).Count(&count).Error

	return count, err
}

func (db *Database) SearchCars(query string) ([]Car, error) {
	var cars []Car

	err := db.Gorm.Model(&Car{}).Where("to_tsvector('simple', id) @@ to_tsquery(?)", query).Find(&cars).Error

	return cars, err
}
