package main

func CheckAll(files []string) {
	waitGroup.Add(3)
	go checkGo(files)
	go checkRust(files)
	go checkYarn(files)
	waitGroup.Wait()
}

func checkRust(files []string) {
	defer waitGroup.Done()
	for _, file := range files {
		if file == "Cargo.toml" {
			Languages = append(Languages, "rust")
		}
	}
}

func checkYarn(files []string) {
	defer waitGroup.Done()
	for _, file := range files {
		if file == "yarn.lock" {
			Languages = append(Languages, "yarn")
		}
	}
}

func checkGo(files []string) {
	defer waitGroup.Done()
	for _, file := range files {
		if file == "go.mod" {
			Languages = append(Languages, "go")
		}
	}
}
