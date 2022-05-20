package main

import (
	"log"
	"net/http"

	"mini.com/controllers"
	"mini.com/dosens"
	"mini.com/tabels"

	"time"

	"github.com/gin-gonic/gin"

	jwt "github.com/appleboy/gin-jwt/v2"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var IdentityKey = "id"

func main() {
	r := gin.Default()
	db := tabels.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	var us tabels.User

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "Toriq Zone",
		Key:         []byte("kbjhfhgcjhbhfcd"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*tabels.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &tabels.User{
				UserName: claims[IdentityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password
			var student tabels.Mahasiswa
			db.Where("nama = ?", userID).Where("password = ?", password).Find(&student)
			us.UserName = student.Nama

			if (userID == student.Nama && password == student.Password) || (userID == "test" && password == "test") {
				return &tabels.User{
					UserName:  userID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*tabels.User); ok && v.UserName == us.UserName {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/historiPresensi", controllers.HistoriPresensi)
		auth.GET("/akumulasiPresensi", controllers.AkumulasiPresensi)
		auth.POST("/presensi", controllers.CreatePresensi)
		auth.GET("/jadwal", controllers.GetJadwal)

		auth.PUT("editJadwal", dosens.EditJadwal)
		auth.GET("akumulasi/:nip", dosens.DosenAkumulasi)
		auth.GET("melihatpresensi/:nip", dosens.MelihatPresensi)
		auth.GET("melihatjadwal/:nip", dosens.LihatJadwal)
		auth.PUT("mengubahakses/:nip", dosens.MengubahAkses)
	}

	if err := http.ListenAndServe(":"+"8080", r); err != nil {
		log.Fatal(err)
	}

	r.Run()
}
