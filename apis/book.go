package apis

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"onego/common"
	"onego/config"
  "onego/int"
  "onego/global"
)

var (
	NovelsRulesMap = config.NovelsRulesMap
)

// SearchAuthors returns all novels resource that you serached
func BookView(c *gin.Context) {
	var currentRule config.NovelRule
	var ok bool
	bookid := c.Param("bookid")
	if bookid != "" {

    log.Println("Bookid request:"+bookid)
 nnovel := &int.Novel{}
 rel, err := global.NB_DB.Where("id = ?", bookid).Get(nnovel)
  
 if !rel || err != nil {
 c.JSON(http.StatusOK, gin.H{"msg": "查询错误"})
 } else {
 c.JSON(http.StatusOK, gin.H{"user": nnovel})
 
  }
}
