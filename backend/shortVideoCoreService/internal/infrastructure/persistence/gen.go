package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	gormdb, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/doutok?parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/infrastructure/persistence/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions

	g.ApplyBasic(
		// Generate struct `User` based on table `users`
		g.GenerateModel("user"),
		g.GenerateModel("video"),
	)
	// Generate the code
	g.Execute()
}
