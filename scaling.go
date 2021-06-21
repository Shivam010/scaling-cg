// Scaling the crispyness with the giggles of togetherness.
package main

import (
	"fmt"

	cg "github.com/Shivam010/crispy-giggle/v3"
	test "github.com/Shivam010/crispy-giggle/v3/testdata"
)

func main() {
	cg.Main(cg.Main)

	cg.License()
	cg.LicenseDir()
	cg.TestData()

	fmt.Println("------")

	fs, err := test.File.ReadDir(".")
	fmt.Println("error", err)

	fmt.Println("------")

	for _, f := range fs {
		// io.ReadAll(f)
		fmt.Println(f.Name())
		fmt.Println(f.IsDir())
		d, err := test.File.ReadFile(f.Name())
		fmt.Println(string(d), err)
		fmt.Println()
	}
}
