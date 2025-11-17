package main

import (
	"flag"
	"latihan-hris/config"
	migration "latihan-hris/database"
	"latihan-hris/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Buat flag untuk migrasi
	runMigration := flag.Bool("migrate", false, "Jalankan migrasi dan seeder awal")
	runDown := flag.Bool("down", false, "Turunkan satu langkah migrasi terakhir")
	dropAll := flag.Bool("drop", false, "Hapus semua tabel di database")
	flag.Parse()

	// load .env
	config.LoadEnv()

	// koneksi database
	config.ConnectDatabase()

	// Jalankan migrasi --migrate
	if *runMigration {
		migration.RunMigration()
		return
	}

	// Rollback satu langkah --down
	if *runDown {
		migration.RunDown()
		return
	}

	// Drop semua tabel --drop
	if *dropAll {
		migration.RunDropAll()
		return
	}

	// inisialisasi route Gin
	r := gin.Default()
	// r.Use(func(c *gin.Context) {
	// 	origin := c.Request.Header.Get("Origin")
	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, ngrok-skip-browser-warning")
	// 	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	// 	// Tangani preflight OPTIONS agar tidak 403
	// 	if c.Request.Method == "OPTIONS" {
	// 		c.AbortWithStatus(204)
	// 		return
	// 	}

	// 	c.Next()
	// })

	// r.Use(cors.Default())
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{
	// 	"http://192.168.0.82:8080", // sesuaikan IP lokal frontend kamu
	// 	"https://33a1e5cecbcc.ngrok-free.app"}, // izinkan semua origin
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "ngrok-skip-browser-warning"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: false,
	// }))

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{
	// 		"http://192.168.0.82:5173",            // frontend di device lain (local IP)
	// 		"http://localhost:5173",               // frontend lokal
	// 		"https://33a1e5cecbcc.ngrok-free.app", // backend via ngrok
	// 		"https://*.ngrok-free.app",            // untuk jaga-jaga subdomain ngrok berubah
	// 	},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "ngrok-skip-browser-warning"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))

	// Tambahkan ini tepat setelah middleware CORS
	// r.OPTIONS("/*path", func(c *gin.Context) {
	// 	c.Status(204)
	// })

	// batas maksimal upload file
	r.MaxMultipartMemory = config.MaxUploadMB << 20

	// static route untuk akses file
	r.Static("/uploads", config.UploadPath)

	// route default
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin!",
		})
	})

	// daftar route
	routes.RegisterRoute(r)
	routes.AuthRoute(r)

	// Jalankan server
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	log.Printf("Server running on port %s", appPort)
	r.Run(":" + appPort)

}
