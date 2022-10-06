package Router
import(
	"blog/Controllers"
	"github.com/gin-gonic/gin"
)

func Start(){
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/asset","./asset")

	e.GET("/", Controllers.ListIndex)

	e.GET("/register", Controllers.GoRegister)
	e.POST("/register", Controllers.RegisterUser)

	e.GET("/login", Controllers.GoLogin)
	e.POST("/login", Controllers.LoginUser)

	e.GET("/post_index", Controllers.GetPostIndex)
	e.POST("/post", Controllers.AddPost)
	e.GET("/post", Controllers.GoAddPost)

	e.GET("/detail/:id", Controllers.GoDetail)



	e.Run()

}
