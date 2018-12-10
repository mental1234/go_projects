// Maps examples (just like inventories in python)
package main

import "fmt"

func main() {
	fmt.Println("Maps example")
	x := make(map[string] int)
	x["Key"] = 20
	fmt.Println(x["Key"])

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":"Hydrogen",
			"state":"gas",
		},
		"He": map[string]string{
			"Name":"Helium",
			"state":"gas",
		},
	}
	if el, ok := elements["H"]; ok {
		fmt.Println(el["name"],el["state"])
	}
}
