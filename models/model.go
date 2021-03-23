package models

import (
	//导入MYSQL数据库驱动，这里使用的是GORM库封装的MYSQL驱动导入包，实际上大家看源码就知道，这里等价于导入github.com/go-sql-driver/mysql
	//这里导入包使用了 _ 前缀代表仅仅是导入这个包，但是我们在代码里面不会直接使用。
	"fmt"

	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	_ "github.com/joho/godotenv/autoload"
)

//定义全局的db对象，我们执行数据库操作主要通过他实现。
var DB *gorm.DB

func Setup() {
	// 声明err变量，下面不能使用:=赋值运算符，否则_db变量会当成局部变量，导致外部无法访问_db变量
	var err error
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	DB, err = gorm.Open(os.Getenv("DB_CONNECTION"), dsn)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	defer DB.DB().Close()
	//设置数据库连接池参数
	DB.DB().SetMaxOpenConns(100) //设置数据库连接池最大连接数
	DB.DB().SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	AutoMigrateAll()
}

func AutoMigrateAll() {
	DB.AutoMigrate()
}
