package Database

func Connect{
	 db := pg.Connect(&pg.Options{
        User: "postgres",
    })
    defer db.Close()

    err := createSchema(db)
    if err != nil {
        panic(err)
    }
}

func createSchema(db *pg.DB) error {
    for _, model := range []interface{}{(*User)(nil), (*Story)(nil)} {
        err := db.CreateTable(model, &orm.CreateTableOptions{
            Temp: true,
        })
        if err != nil {
            return err
        }
    }
    return nil
}