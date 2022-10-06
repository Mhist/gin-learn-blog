package Dao

import (
	"blog/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Manager interface {
		Register(user *Models.User)
		Login(username string) Models.User
		// 博客操作
		AddPost(post *Models.Post)
		GetAllPost() []Models.Post
		GetPost(pid int) Models.Post

	}


type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init()  {
	dsn := "root:admin123@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatal("Failed to into db:",err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&Models.User{})
}

func (mgr *manager) Register(user *Models.User)  {
	mgr.db.Create(user)
}

func (mgr *manager) Login(username string) Models.User {
	var user Models.User
	mgr.db.Where("username=?",username).First(&user)
	return user
}


// 博客操作
func (mgr *manager) AddPost(post *Models.Post) {
	mgr.db.Create(post)
}
func (mgr *manager) GetAllPost() []Models.Post{
	var posts =make([]Models.Post,10)
	mgr.db.Find(&posts)
	return posts
}

// 博客操作
func (mgr *manager) GetPost(pid int) Models.Post {
	var post Models.Post
	mgr.db.Find(&post,pid)
	return post
}

// 去详情页


