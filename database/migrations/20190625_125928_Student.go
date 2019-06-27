package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Student_20190625_125928 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Student_20190625_125928{}
	m.Created = "20190625_125928"

	migration.Register("Student_20190625_125928", m)
}

// Run the migrations
func (m *Student_20190625_125928) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE Student(`id` int(11) NOT NULL AUTO_INCREMENT,`cpf` varchar(128) NOT NULL,`name` varchar(128) NOT NULL,`email` varchar(128) NOT NULL,`phone` varchar(128) NOT NULL,`birth_date` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Student_20190625_125928) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `Student`")
}
