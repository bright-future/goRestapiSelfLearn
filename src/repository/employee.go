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
	// fmt.Println(redisConn)
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


func InsertEmployee(dbname string, newEmployee models.Employee) (bool, error){
	redisConn := driver.GetConnections()
	employeeByte, _  := json.Marshal(newEmployee)
	employeeStr := string(employeeByte)
	var val bool
	_, err := redis.String(redisConn.Do("HSET", dbname, newEmployee.EmployeeID, employeeStr))
	if (err != nil) {
		fmt.Println(err, dbname, newEmployee, employeeStr)
		return val, err
	}
	val = true
	return val, nil
}


func DeleteEmployee(dbname, id string) (bool, error) {
	redisConn := driver.GetConnections()
	_, err := redis.String(redisConn.Do("HDEL", dbname, id))
	if (err != nil) {
		return false, nil
	}
	return true, nil
}
