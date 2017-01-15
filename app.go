package main

import(
  "fmt"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  _"github.com/julienschmidt/httprouter"

  "github.com/ywoonder/sshop/src"

     "net/http"
)
// grace "gopkg.in/paytm/grace.v1"


func printHeader(){

  header:=`


 .d8888b.   .d8888b.  888    888  .d88888b.  8888888b.
d88P  Y88b d88P  Y88b 888    888 d88P" "Y88b 888   Y88b
Y88b.      Y88b.      888    888 888     888 888    888
 "Y888b.    "Y888b.   8888888888 888     888 888   d88P
    "Y88b.     "Y88b. 888    888 888     888 8888888P"
      "888       "888 888    888 888     888 888
Y88b  d88P Y88b  d88P 888    888 Y88b. .d88P 888
 "Y8888P"   "Y8888P"  888    888  "Y88888P"  888


`
fmt.Printf("%s",header)

}
var Master *sql.DB
func openDB(){
  println("init database")
  db, err := sql.Open("mysql","root:@tcp(localhost:3306)/sshop")
  if err!=nil{
    fmt.Println("err ",err)
  }
  Master = db
  err = db.Ping()
}

func setRoutes(){
  println("set routes")
// julienschmidt router
//  router := httprouter.New()
//  router.GET("/index",handlerIndex )
//  log.Fatal(grace.Serve("8080", router))

  //nethttp router
  http.HandleFunc("/", handlerIndex)
  http.ListenAndServe(":8080", nil)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
fmt.Fprintf(w, "HELLO WORLD! SSHOP HERE~" )
}

func main() {
    openDB()
//    setRoutes()

    println("test insert product")

    testInsertProduct(Master)
printHeader()
    http.HandleFunc("/", handlerIndex)
    http.ListenAndServe(":9000", nil)

}
