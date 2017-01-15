package main
import(
  "class"
  "database/sql"
)

func TestInsertProduct(db *sql.DB){

  p := Product{
    Name : "Yunita",
    Description : "Hello world",
    Price : 123456,
    Stock : 7,
  }

  err := class.InsertProduct(db)
  println("error: "err)
  return err
}
