package dict

import "fmt"

func MapOperations() {
	// var m map[string]string

	m := make(map[string]string)

	m["name"] = "Muhammad"
	// name := m["name"]
	m["job"] = "Messanger of Allah"
	m["calc"] = "5 + 5 = 10"

	iterateMap(&m)
	delete(m,"calc")
	println("After DELETE of calc")

	iterateMap(&m)

	println("CHECK if key exists")
	getMapKey(&m)

	println("POPULATE map")
	populateMapWithInitValues()

}


func getMapKey(m *map[string]string){
	_, ok := (*m)["abc"]
	fmt.Println(ok)
}

func iterateMap(m *map[string]string){
	for key, value := range *m {
		fmt.Printf("%s \t---->\t %s\n", key, value)
	}

}

func populateMapWithInitValues(){
	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	for k, v := range commits{
		fmt.Printf("%v \t---> %v\n", k, v)
	}
}
