things := []string{"maleta", "bolso", "boleto", "vaso", "casa"}

	fmt.Println(things)

	fmt.Println("First for loop: ")
	fmt.Println("for i:=0;  i < 10; i++{\n fmt.Println(i + start)\n}")
	for i:=0; i < 10; i++{
		fmt.Println(i + start)
	}

	fmt.Println("\nSecond for loop: ")
	fmt.Println("for i:=0; i<len(things); i++{\n fmt.Println(things[i])")
	for i:=0; i<len(things); i++{
		fmt.Println(things[i])
	}

	fmt.Println("\nThird for loop: ")
	fmt.Println("for i:=range things {\n fmt.Println(things[i])")
	for i:=range things {
		fmt.Println(things[i])
	}

	fmt.Println("\nFourth for loop & break:")
	fmt.Println("for start < 100 {\nstart += start\nif start == 32 {\nbreak\n}\nfmt.Println(\"START is now: \", start)")
	for start < 100 {
		start += start
		if start == 32 {
			break
		}
		fmt.Println("START is now: ", start)
	}