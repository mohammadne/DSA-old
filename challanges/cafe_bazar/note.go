package main

// 	var (
// 		id int
// 		name string
// 		from string
// 		to string
// 		through []int
// 		speedLimit int
// 		speedLimit int
// 		‫‪biDirectional‬‬ int
// 		length‬‬ int
// 	)

// 	input, _ := reader.ReadString('\n')
// 		input = strings.Trim(input, "\n")
// 		input = strings.Trim(input, "[")
// 		input = strings.Trim(input, "]")
// 		values := strings.Split(input, ",")

// 		through := make([]int, 0, 10)
// 		for _, value := range values {
// 			number, _ := strconv.Atoi(value)
// 			through = append(through, number)
// 		}

// 		fmt.Println("speed_limit=?")
// 		fmt.Scan(&speedLimit)

// 		fmt.Println("‫‪length‬‬=?")
// 		fmt.Scan(&‫‪length‬‬)

// 		// var ‫‪length‬‬ int
// 		// fmt.Println("‫‪length‬‬=?")
// 		// fmt.Scan(‫‪&length‬‬)

// 		// var ‫‪biDirectional‬‬ int
// 		// fmt.Println("‫‪bi_directional‬‬=?")
// 		// fmt.Scan(‫‪&biDirectional‬‬)

// 		road := Road{
// 			id:            id,
// 			name:          name,
// 			from:          from,
// 			to:            to,
// 			through:       through,
// 			speedLimit:    speedLimit,
// 			length:        length,
// 			biDirectional: biDirectional,
// 		}

// 		roads.add(road)

// 		fmt.Printf("Road with id%d added!", id)

// 		action := printOptions(
// 			"Select your next action",
// 			[]string{"Add another Road", "Main Menu"},
// 		)

// 		switch action {
// 		case 1:
// 			processAdd(cities, roads, 2)
// 		case 2:
// 			processMainMenue(cities, roads)

// 		}
// }

// v := reflect.ValueOf(city)

// for i := 0; i < v.NumField(); i++ {

// 	var input interface{}
// 	fmt.Scan(input)

// 	valuePtr := reflect.ValueOf(input)
// 	v.Field(i).Set(valuePtr)

// t := v.Field(i).Type()
// var size t
// fmt.Scan(&size)
// // a := v.Field(i).Type()

// values[i] = v.Field(i).Interface()
// v.FieldByName()
// }

// func processAdd(models models) {
// 	modelsTypeReflect := reflect.TypeOf(models).Elem()
// modelsTypeStr := fmt.Sprintf("%s", modelsTypeReflect)
// modelsType := strings.Split(modelsTypeStr, ".")

// modelPtr := reflect.New(modelsTypeReflect)

// fmt.Println(intPtr2)

// v := reflect.ValueOf(modelPtr)

// values := make([]interface{}, v.NumField())

// fmt.Println(v.NumField())
// for i := 0; i < v.NumField(); i++ {
// 	fmt.Println(v.Field(i).Interface)
// }

// fmt.Println(values)

// fmt.Printf("%s with id=%d added!", name, 0)

// value := modelsTypeReflect{}

// v := reflect.ValueOf(models)

// values := make([]interface{}, v.NumField())

// for i := 0; i < v.NumField(); i++ {
// 	values[i] = v.Field(i).Interface()
// }

// fmt.Println(values)
// }

// func processAdd(reader *bufio.Reader, m models) {
// 	if model == 0 {
// 		model = getModel(reader)
// 	}

// 	switch model {
// 	case 1:
// 		var id int
// 		fmt.Scan(&id)

// 		var name string
// 		fmt.Scan(&name)

// 		city := city{
// 			id:   id,
// 			name: name,
// 		}

// 		cities.add(city)

// 		fmt.Printf("City with id=%d added!\n", id)

// 		number := printOptions(
// 			reader,
// 			"Select your next action",
// 			[]string{"Add another City", "Main manu"},
// 		)

// 		switch number {
// 		case 1:
// 			return processAdd(reader, 1)
// 		}
// 	}
// }
