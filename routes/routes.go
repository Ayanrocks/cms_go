package routes

import (
	"github.com/cms/Database"
	"github.com/cms/types"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

//Route file
func Routes(App *iris.Application) {
	App.Handle("GET", "/", func(ctx iris.Context) {

		ctx.HTML("<h1>Welcome to IRIS NEW</h1>")
	})

	App.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Pong!"})
	})

	App.Post("/posts/create", func(ctx iris.Context) {
		var data types.Post
		err := ctx.ReadJSON(&data)
		if err != nil {
			golog.Error("PARSE ", err)
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}

		err = Database.Insert(data)

		if err != nil {
			golog.Error("ENTER SYCC ", err)
			ctx.StatusCode(500)
			ctx.WriteString(err.Error())

		}

		ctx.JSONP(map[string]string{"msg": "SUCCESS"})
	})

}
