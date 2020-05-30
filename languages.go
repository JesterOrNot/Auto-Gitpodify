package main

func CheckAll(files []string) {
	waitGroup.Add(5)
	go makeChecker("go.mod", "go")(files)
	go makeChecker("Cargo.toml", "rust")(files)
	go makeChecker("yarn.lock", "yarn")(files)
	go makeChecker("package.json", "npm")(files)
	go makeChecker("Makefile", "make")(files)
	waitGroup.Wait()
}

func makeChecker(targetFileName string, language string) func(files []string) {
	return func(files []string) {
		defer waitGroup.Done()
		for _, file := range files {
			if file == targetFileName {
				Languages = append(Languages, language)
			}
		}
	}
}
