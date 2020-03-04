package repository

import(
	"../driver"
	"../models"
	"fmt"
	"log"
	"github.com/gomodule/redigo/redis"
	"encoding/json"
	// "../models"
)


func GetByID(dbname, id string) (models.Employee, error){
	redisConn := driver.GetConnections()
	fmt.Println(redisConn)
	var employee models.Employee;
	val, err := redis.String(redisConn.Do("HGET", dbname, id))
	if (err != nil) {
		log.Println(err)
		return employee, err
	} 
	err = json.Unmarshal([]byte(val), &employee)
	if (err != nil) {
		log.Println(err)
		return employee, err
	}
	// fmt.Println(val, employee)
	return employee, nil
}