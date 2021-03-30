package int
import (
 _ "github.com/go-sql-driver/mysql"
 "github.com/go-xorm/xorm"
 "log"
)
  
type Novel struct {
 Id int `xorm:"INT(10) 'id'"` 
 Name string `xorm:"VARCHAR(20) 'name'"`
 Cover string `xorm:"VARCHAR(200) 'cover'"` 
 Intro int64 `xorm:"text 'intro'"` 
 Author_id int `xorm:"INT(11) 'author_id'"` 
 Is_over int `xorm:"TINYINT(4) 'is_over'"` 
 Source_id int `xorm:"int(11) 'source_id'"` 
 Last_update int `xorm:"int(10) 'last_update'"` 
 Last_chapter_id	 int `xorm:"int(11) 'last_chapter_id'"` 
 Last_chapter_name int `xorm:"varchar(100) 'last_chapter_name'"` 
 Channel_id int `xorm:"smallint(5) 'channel_id'"` 
 Category_id int `xorm:"smallint(5) 'category_id'"` 
 Num_allvisit int `xorm:"int(10) 'num_allvisit'"` 
 Num_monthvisit int `xorm:"int(10) 'num_monthvisit'"` 
 Num_weekvisit int `xorm:"int(10) 'num_weekvisit'"` 
 Num_dayvisit int `xorm:"int(10) 'num_dayvisit'"` 
 Status int `xorm:"int(10) 'status'"` 

}

//定义orm引擎
var x *xorm.Engine 
//创建orm引擎
func Mysql() {
 var err error
 x, err = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/xorm?charset=utf8mb4")
 if err != nil {
 log.Fatal("数据库连接失败:", err)
 }
 if err := x.Sync(new(Novel)); err != nil {
 log.Fatal("数据表同步失败:", err)
 }
}
