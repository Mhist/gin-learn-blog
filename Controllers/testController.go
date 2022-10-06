package Controllers
import(
	"blog/Dao"
	"blog/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)
func AddUser(c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := Models.User{
		Username: username,
		Password: password,
	}
	Dao.Mgr.Register(&user)
}


func ListIndex(c *gin.Context)  {
	c.HTML(200,"index.html",nil)
}

func ListRegister(c *gin.Context)  {
	c.HTML(200,"register.html",nil)
}


func ListLogin(c *gin.Context)  {
	c.HTML(200,"login.html",nil)
}
// 注册 输入账号密码
func RegisterUser(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := Models.User{
		Username: username,
		Password: password,
	}
	Dao.Mgr.Register(&user)
	c.Redirect(301,"/")
}
// 去注册页面
func GoRegister(c *gin.Context){
	c.HTML(200,"register.html",nil)
}
// 去登录页面
func GoLogin(c *gin.Context){
	c.HTML(200,"login.html",nil)
}
// 登录 账号密码 接口
func LoginUser(c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)
	u := Dao.Mgr.Login(username)
	if u.Username == ""{
		c.HTML(200,"login.html","用户名不存在")
		fmt.Println("用户名不存在")
	}else{
		if u.Password != password{
			fmt.Println("密码错误")
			c.HTML(200,"login.html","密码错误")
		}else{
			fmt.Println("登录成功")
			c.Redirect(301,"/")
		}
	}

}

func GetPostIndex(c *gin.Context){
	posts := Dao.Mgr.GetAllPost()
	c.HTML(200,"postIndex.html",posts)
}

func AddPost(c *gin.Context){
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")
	post := Models.Post{
		Title: title,
		Tag: tag,
		Content: content,
	}
	Dao.Mgr.AddPost(&post)
	c.Redirect(301,"post_index")
}

func GoAddPost(c *gin.Context){
	c.HTML(200,"post.html",nil)
}

func GoDetail(c *gin.Context){
	// 获取路径传参
	id := c.Param("id")
	id2, _ := strconv.Atoi(id)
	post := Dao.Mgr.GetPost(id2)
	c.HTML(200,"detail.html",post)
}





