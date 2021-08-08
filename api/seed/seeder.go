package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/prashantvnsi/myblog/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Prashant Tiwari",
		Email:    "prashant.tiwari@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Roger Federer",
		Email:    "roger.federer@gmail.com",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Saving the World",
		Content: "Everyone should think about it",
	},
	models.Post{
		Title:   "How to become best",
		Content: "Hard Work, Priority",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
