package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run injector.go executable1 executable2")
		return
	}

	bin1, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	bin2, err := os.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("tmpfile.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(combinedCode(bin1, bin2))
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("go", "build", "-o", os.Args[1], "tmpfile.go")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	os.Remove("tmpfile.go")
}

func combinedCode(bin1, bin2 []byte) string {
	return fmt.Sprintf(`
	package main

	import (
		"log"
		"os"
		"os/exec"
	)


	func main() {
		err := runBinary(%#v)
		if err != nil {
			log.Fatal(err)
		}
		err = runBinary(%#v)
		if err != nil {
			log.Fatal(err)
		}
	}
	
	func runBinary(data []byte) error {
		os.Remove("tmpexec")
		err := os.WriteFile("tmpexec", data, 0755)
		if err != nil {
			return err
		}

		cmd := exec.Command("./tmpexec")
    	cmd.Stdout = os.Stdout
    	cmd.Stderr = os.Stderr

    	err = cmd.Run()
    	if err != nil {
        	return err
    	}
		os.Remove("tmpexec")
		return nil
	}
	`, bin1, bin2)
}
