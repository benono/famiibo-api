package router

import (
	"go-rest-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, pc controller.IPayeeController, ac controller.IAccountController, cc controller.ICategoryController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:3001", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true, // Cookieの送受信を可能にする
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode, POSTMAN確認用
		// CookieMaxAge: 60,
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	t := e.Group("/tasks")
	// ミドルウェアを追加
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token", // jwtTokenの格納場所（CookieにTokenという名前で保存）
	}))

	// Accounts
	a := e.Group("/accounts")
	a.GET("", ac.FindAll)
	a.GET("/:accountId", ac.FindById)
	a.POST("", ac.Create)
	a.PUT("/:accountId", ac.Update)
	a.DELETE("/:accountId", ac.Delete)

	// Payees
	p := e.Group("/payees")
	p.GET("", pc.FindAll)
	p.GET("/:payeeId", pc.FindById)
	p.POST("", pc.Create)
	p.PUT("/:payeeId", pc.Update)
	p.DELETE("/:payeeId", pc.Delete)

	// Categories
	c := e.Group("/categories")
	c.GET("", cc.FindAll)
	c.GET("/:categoryId", cc.FindById)
	c.POST("", cc.Create)
	c.PUT("/:categoryId", cc.Update)
	c.DELETE("/:categoryId", cc.Delete)

	// Stores
	// s := e.Group("/stores")
	// s.GET("", sc.FindAll)
	// s.GET("/:storeId", sc.FindById)
	// s.POST("", sc.Create)
	// s.PUT("/:storeId", sc.Update)
	// s.DELETE("/:storeId", sc.Delete)
	return e
}
