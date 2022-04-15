package main

import "gorm.io/gorm"

type Employees struct {
	gorm.Model
	Position  string  `json:"pos"`
	EmpName   string  `json:"name"`
	EmpSalary float64 `json:"salary"`
	Email     string  `json:"email"`
}

//start
