package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Instructors_20190625_131153 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Instructors_20190625_131153{}
	m.Created = "20190625_131153"

	migration.Register("Instructors_20190625_131153", m)
}

// Run the migrations
func (m *Instructors_20190625_131153) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE Instructors(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`email` varchar(128) NOT NULL,`hour_price` float NOT NULL,`certificates` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Instructors_20190625_131153) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `Instructors`")
}
