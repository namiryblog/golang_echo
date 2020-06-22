package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
  )

  func main() {
	// エコーのインスタンス→実態みたいなもの  
	echo := echo.New()

	// エコーノミドルウェア
	//　ミドルウェアとは、処理が行われる前とか後に動くもの・・・？？
	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	// 実際にリクエストするＵＲＬ　！！
	// "/"にアクセスすると呼べる
	echo.GET("/", hello)

	// サーバーをスタートさせるコード
	//　localhost:3000 でサーバーが動きます！！
	echo.Logger.Fatal(echo.Start(":30000"))
  }

  func hello(c echo.Context) error {
	  // hello が呼ばれたときにstatus ok だったら"Hello, Namiry !?"がでるんかな？？
	  return c.String(http.StatusOK, "Hello, Namiry !?")
  }

  