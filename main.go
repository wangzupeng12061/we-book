package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	redis1 "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/wangzupeng12061/we-book/internal/pkg/ginx/middlewares/ratelimit"
	"github.com/wangzupeng12061/we-book/internal/repository"
	"github.com/wangzupeng12061/we-book/internal/repository/dao"
	"github.com/wangzupeng12061/we-book/internal/service"
	"github.com/wangzupeng12061/we-book/internal/web"
	"github.com/wangzupeng12061/we-book/internal/web/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	db := initDB()
	u := initUser(db)
	server := initWebServer()
	u.RegisterRoutes(server)
	server.Run(":8080")
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	server.Use(ratelimit.NewBuilder(redisClient, time.Second, 100).Build())
	//server.Use(redislimit.NewRedisActiveLimit(redisClient,3,"",)
	server.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"http://localhost:3000"},
		//AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"authorization", "content-type"},
		ExposeHeaders:    []string{"x-jwt-token"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "example.com")
		},
		MaxAge: 12 * time.Hour,
	}))
	//store := cookie.NewStore([]byte("secret"))
	//store := memstore.NewStore([]byte("0sbBgcdb79Dd4fbjBF336s855sA6b20w"),
	//	[]byte("es982F0250Bnt1qExlF12l631nmEge1y"))
	store, err := redis1.NewStore(16, "tcp", "localhost:6379", "",
		[]byte("0sbBgcdb79Dd4fbjBF336s855sA6b20w"),
		[]byte("es982F0250Bnt1qExlF12l631nmEge1y"))
	if err != nil {
		panic(err)
	}
	server.Use(sessions.Sessions("mysession", store))
	//server.Use(middleware.NewLoginMiddlewareBuilder().IgnorePaths("/users/login").Build())
	server.Use(middleware.NewLoginJWTMiddlewareBuilder().IgnorePaths("/users/login").Build())
	return server
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}
func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}

	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}
