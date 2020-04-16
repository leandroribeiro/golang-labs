package seed

import (
	"github.com/jinzhu/gorm"
	"github.com/leandroribeiro/go-labs/api-lab3/api/models"
	"github.com/prometheus/common/log"
)

var users = []models.User{
	models.User{
		Nickname: "Obiwan Kenobi",
		Email:    "obiwan@jedi.org",
		Password: "password",
	},
	models.User{
		Nickname: "Anakin Skywalker",
		Email:    "anakin@jedi.org",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Episode I",
		Content: "Episode I content",
	},
	models.Post{
		Title:   "Eposode II",
		Content: "Episode II content",
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
