package app

 
import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "sync"
 
    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)
 
var (
    db   *sql.DB
    once sync.Once
)
 
func GetInstance() *sql.DB {
    
    once.Do(func() {
        err := godotenv.Load()
        if err != nil {
            log.Fatalf("Error loading .env file:%v", err)
        }
        user := os.Getenv("DB_USER")
        password := os.Getenv("DB_PASSWORD")
        host := os.Getenv("DB_HOST")
        port := os.Getenv("DB_PORT")
        dbname := os.Getenv("DB_NAME")
 
        dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
        db, err = sql.Open("mysql", dsn)
        if err != nil {
            log.Fatalf("Failed to open DB: %v", err)
        }
 
        if err := db.Ping(); err != nil {
            log.Fatalf("Failed to connect to DB : %v", err)
        }
 
        log.Println("Database Connection Established")
 
    })
    return db
}