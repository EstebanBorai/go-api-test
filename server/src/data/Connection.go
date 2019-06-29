package data

import (
	"github.com/estebanborai/songs-share-server/server/src/helpers/gimlet"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var connectionString = "root:root@tcp(songs-share-db)/songs-share?charset=utf8mb4&parseTime=True&loc=Local"

// NOTE: docker port <container id> 3306
var localConnectionString = "root:root@tcp(0.0.0.0:3306)/songs-share?charset=utf8mb4&parseTime=True&loc=Local"

// Connection Creates a connection to the database and returns it
func Connection(c *gin.Context) (db *gorm.DB) {
	db, err := gorm.Open("mysql", localConnectionString)

	if err != nil {
		gimlet.InternalServerError(c, "Unable to connect to the database")
	}

	return db
}
