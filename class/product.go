package class

import(
  "fmt"
  "database/sql"
)

type(

  Product struct{
    ID int64 `db:"id" json:"product_id"`
    Name string `db:"name" json:"product_name"`
    Description string `db:"description" json:"product_desc"`
    Price int64 `db:"price" json:"product_price"`
    Stock int64 `db:"stock" json:"product_stock"`
    CreateTime string `db:"create_time" json:"product_create_time"`
    ImgUrl string
  }

  SearchFilter struct{
    LowestPrice int64
    HighestPrice int64
    CreateTime string
    Name string
    Description string
  }
)


func InsertProduct(p Product, *sql.DB)error{
    query := fmt.Sprintf(`insert into
                products
                (name,description, price,stock, create_time)
              values
                (%s,%s, %d, %d, now())
              `, p.Name, p.Description, p.Price, p.Stock, p.CreateTime)
    println(query)
    return nil
}

func GetProductByID(ID int64)Product{
  return Product{}
}

func DeleteProduct(ID int64)error{
  return nil
}
