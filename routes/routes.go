package routes

import "github.com/kataras/iris"

//Route file
func Routes(App *iris.Application) {
	App.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome to IRIS NEW</h1>")
	})

	App.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Pong!"})
	})

	App.Post("/posts/create", func(ctx iris.Context) {

	})

}
