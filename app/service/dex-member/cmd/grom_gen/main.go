package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./app/service/dex-member/storage/mysql/query",
		ModelPkgPath:      "./app/service/dex-member/storage/mysql/model",
		WithUnitTest:      true,
		FieldCoverable:    true,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})
	dns := "root:@tcp(127.0.0.1:3306)/dex?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dns))
	g.UseDB(db)
	g.ApplyBasic(g.GenerateModel("user"))
	g.Execute()
}
