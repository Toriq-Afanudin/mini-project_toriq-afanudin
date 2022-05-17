package main

import (
	"log"

	"mini.com/dosens"
	"mini.com/mahasiswas"
	"mini.com/tabels"

	"time"

	"github.com/gin-gonic/gin"

	jwt "github.com/appleboy/gin-jwt/v2"
)

type login struct {
	Nama string `json:"nama"`
	Nim  string `json:"nim"`
}

var identityKey = "id"

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func main() {
	r := gin.Default()

	db := tabels.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	var nama string

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "Toriq Zone",
		Key:         []byte("njksdendjnjhdbjsnddjhjbdgj"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Nama
			password := loginVals.Nim

			var mahasiswa tabels.Mahasiswa
			db.Where("nama = ?", userID).Where("nim = ?", password).Find(&mahasiswa)
			nama = mahasiswa.Nama

			if (userID == mahasiswa.Nama && password == mahasiswa.Nim) || (userID == "test" && password == "test") {
				return &User{
					UserName:  userID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == nama {
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
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

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
		auth.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "berhasil masuk dengan token jwt"})
		})
		auth.PUT("editJadwal/:nip ", dosens.EditJadwal)
		auth.GET("akumulasi/:nip", dosens.DosenAkumulasi)
		auth.GET("melihatpresensi/:nip", dosens.MelihatPresensi)
		auth.GET("melihatjadwal/:nip", dosens.LihatJadwal)
		auth.PUT("mengubahakses/:nip", dosens.MengubahAkses)

		auth.POST("presensi/:nim", mahasiswas.Presensi)
		auth.GET("akumulasiMahasiswa/:nim", mahasiswas.MahasiswaAkumulasi)
		auth.GET("presensi/:nim", mahasiswas.LihatPresensi)
	}

	r.Run()
}
