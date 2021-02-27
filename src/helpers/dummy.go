package helpers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"myurl.com/inventory/models"
)

func InsertDummyData(db *gorm.DB) {

	i := 0
	for i < 5 {

		society := models.Society{
			Email:    "testemail" + fmt.Sprintf("%d", i) + "@test.com",
			Password: "testpass",
			Details:  "This is a dummy society",
		}

		// Add Society
		db.Create(&society)

		// Items
		j := 0
		for j < 5 {
			item := models.Item{
				Details:   "This is a test Item" + fmt.Sprintf("%d", j),
				Quantity:  420,
				Available: 69,
				SocietyId: society.UUID,
			}

			db.Create(&item)
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
