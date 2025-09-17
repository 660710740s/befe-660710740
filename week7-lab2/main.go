package main

import (
	"fmt"
	"os"
)
func getEnv(key,DefaultValue string) string {
	if value := os.Getenv("DB_HOST");value!=""{
		return value
	}
	return DefaultValue
}
func main (){
	host := getEnv("DB_HOST","")
	name := getEnv("DB_NAME","")
	user := getEnv("DB_USER","")
	password :=getEnv("DB_PASSWORD","")
	port := getEnv("DB_PORT","")

	const := fmt.Sprint("host=%s port=%s user=%s password=%s dbname=%s",host, port, user, password, name)
	fmt.Println(const)	
}