package main

import (
	"contacts/database"
	"contacts/handlers"
	"flag"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

var (
	PORT string
	DSN  string
)

func main() {
	PORT = os.Getenv("PORT")
	DSN = os.Getenv("DB_CONN")
	if PORT == "" {
		flag.StringVar(&PORT, "port", "8084", "--port=8084 or --port 8084 or -port=8084")
	}
	if DSN == "" {
		flag.StringVar(&DSN, "dsn", "host=localhost user=admin password=admin123 dbname=contactsdb port=65432 sslmode=disable TimeZone=Asia/Shanghai", "--dsn=host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai")
	}
	flag.Parse()

	db, err := database.GetConnection(DSN)
	if err != nil {
		log.Fatal(err)
	}

	//gin.SetMode(gin.ReleaseMode)
	//gin.DisableConsoleColor()
	router := gin.Default()
	//log.Println("server has started and listening on port", PORT)

	router.GET("/ping", handlers.Ping)

	chandler := new(handlers.Contact)
	cdb := new(database.ContactDB)
	cdb.DB = db
	chandler.IContact = cdb

	router.POST("/contact", chandler.Create)
	router.GET("/contact/:id", chandler.Get)

	if err := router.Run(":" + PORT); err != nil {
		log.Println(err)
		runtime.Goexit()
	}

	// http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("pong"))
	// 	w.WriteHeader(http.StatusOK)
	// })

	// http.HandleFunc("/health", health)
	// log.Println("server has started and listening on port", PORT)
	// http.ListenAndServe(":"+PORT, nil) // :8084 means all network interfaces on this machine can be used
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
	w.WriteHeader(200)
}

// wifi modem
// lan
// bluetooth
// thuderbolt

// lo loop back
// software defined networks
// bridge
// flannel
