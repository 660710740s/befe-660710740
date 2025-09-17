package main

import (
	"fmt"
	"os"
	"database/sql"
	_"github.com/lib/pq"
	"log"
)
func getEnv(key,DefaultValue string) string {
	if value := os.Getenv("DB_HOST");value!=""{
		return value
	}
	return DefaultValue
}
var db * sql.DB

func initDB(){
	var err error 
	host := getEnv("DB_HOST","")
	name := getEnv("DB_NAME","")
	user := getEnv("DB_USER","")
	password :=getEnv("DB_PASSWORD","")
	port := getEnv("DB_PORT","")

	conSt := fmt.Sprint("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, name)
	
	db,err=sql.Open("postgres",conSt)
	//fmt.Println(conSt)	
	if err!=nil{
		log.Fatal("failed to open database")

	}
	err=db.Ping()
	if err!=nil{
		log.Fatal("failed to connect to database")
	}
	log.Println("successfully connected to database")

}
func main (){
	initDB()
}