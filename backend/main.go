package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Yamako76/go-react/infra"
	"github.com/Yamako76/go-react/infra/mysql"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	db, err := infra.NewDB()
	if err != nil {
		log.Fatalf("データベースの接続に失敗しました: %v", err)
	}
	userRepository := mysql.NewUserRepository(db)

	RegisterRoutes(e, Registry{
		UserRepository: userRepository,
	})

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatalf("サーバーにエラーが発生しました: %v", err)
	}
}

// func secureEndpoint(c echo.Context) error {
// 	return c.String(http.StatusOK, "This is a secured endpoint!")
// }

// func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// JWTの検証を行う処理をここに記述
// 		// トークンが無効な場合、Unauthorizedエラーを返す
// 		return next(c)
// 	}
// }
