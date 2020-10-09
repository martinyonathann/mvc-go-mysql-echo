package models

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"mvc-go-mysql-echo/db"
)

type Employee struct {
	Id     string `json:"id"`
	Name   string `json:"employee_name"`
	Salary string `json:"employee_salary"`
	Age    string `json:"employee_age"`
}

type Employees struct {
	Employees []Employee `json:"employee"`
}

var con *sql.DB

func GetEmployee() Employees {
	con := db.CreateCon()
	sqlStatement := "Select id, employee_name, employee_age, employee_salary FROM employee order by id"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := Employees{}

	for rows.Next() {
		employee := Employee{}
		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)
		if err2 != nil {
			fmt.Println(err2)
		}
		result.Employees = append(result.Employees, employee)
	}
	return result
}
