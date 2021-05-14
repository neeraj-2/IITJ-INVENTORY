package helpers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"myurl.com/inventory/models"
)

func InsertDummyData(db *gorm.DB) {

	arr := [5] string{"Robotics Society", "PCLUB", "Drama Society", "Music Society", "Dance Society"}

	i := 0
	for i < 5 {

		society := models.Society{
			Name: arr[i],
			Email:    "testemail" + fmt.Sprintf("%d", i) + "@test.com",
			Password: "testpass",
			Details:  "This is our" + arr[i],
		}

		// Add Society
		db.Create(&society)

		items := [5] string{arr[i]+"Item1", arr[i]+"Item2", arr[i]+"Item3", arr[i]+"Item4", arr[i]+"Item5"}

		// Items
		j := 0
		for j < 5 {
			item := models.Item{
				SocietyName: arr[i],
				Name: items[j],
				Details:   "This is a"+items[j] + "of"+ arr[i],
				Quantity:  420,
				Available: 69,
				SocietyId: society.UUID,
			}

			db.Create(&item)
			// db.Model(&society).Association("Items").Append(&item)
			j += 1
		}

		user := models.User{
			Name:  "dummyUser" + fmt.Sprintf("%d", i),
			Email: "testemailuser" + fmt.Sprintf("%d", i) + "@test.com",
		}
		db.Create(&user)

		i += 1
	}

	user := models.User{
		Name:     "superdummyuser",
		Password: "adminpass",
		Email:    "testemailuser@test.com",
		IsAdmin:  true,
	}
	db.Create(&user)

}
