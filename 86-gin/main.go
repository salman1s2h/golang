package main

import (
	"demo/database"
	"demo/handlers"
	"flag"
	"log"
	"os"

	"math/rand"

	"github.com/gin-gonic/gin"
)

var (
	PORT    string
	DB_CONN string
	// Just to explain main is not the starting point.. This below one is nothing to do with this application
	S int = func() int {
		n := rand.Intn(100)
		return n * n
	}()
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Println("------->", S)
}

func main() {
	e := gin.Default()
	//cargs := os.Args
	//fmt.Println(cargs)
	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err.Error())
	}
	log.SetOutput(f)
	defer f.Close()

	PORT = os.Getenv("PORT")
	if PORT == "" {
		flag.StringVar(&PORT, "port", "8084", "--port=8085 | -port=8085 | --port 8085")
	}
	DB_CONN = os.Getenv("DB_CONN")
	if DB_CONN == "" {
		flag.StringVar(&DB_CONN, "db", `host=localhost user=admin password=admin123 dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai`, "--db=db connection string | -db=8085 | --db 8085")
	}
	flag.Parse()

	//DB_CONN = `host=localhost user=admin password=admin123 dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai`
	db, err := database.GetConnection(DB_CONN)
	if err != nil {
		log.Fatalln(err.Error())
	}

	icontactdb := database.NewContactDB(db)

	e.GET("/", handlers.Hello)
	e.GET("/ping", handlers.Ping)
	e.GET("/health", handlers.Health)
	e.POST("/contact", handlers.NewContacthandler(icontactdb).Create)
	e.GET("/contact/all", handlers.NewContacthandler(icontactdb).GetAll)
	e.DELETE("/contact/:id", handlers.NewContacthandler(icontactdb).DeleteByID)
	e.GET("/contact/:id", handlers.NewContacthandler(icontactdb).GetByID)

	if err := e.Run(":" + PORT); err != nil {
		log.Fatal(err.Error())
	}

}
