package main

import (
	"fmt"

	"github.com/nickmelekh/home_work_basic/hw02_fix_app/printer"
	"github.com/nickmelekh/home_work_basic/hw02_fix_app/reader"
	"github.com/nickmelekh/home_work_basic/hw02_fix_app/types"
)

func main() {
	var path string
	var err error
	var staff []types.Employee

	fmt.Printf("Enter data file path: ")
	_, err = fmt.Scanln(&path)
	if err != nil {
		return
	}

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
