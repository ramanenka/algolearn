package graphs

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Dotviz visializes data in dot format and saves it to [name].png
func Dotviz(name, data string) {
	cmd := exec.Command("dot", "-Tpng")

	f, err := os.Create(fmt.Sprintf("%s.png", name))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cmd.Stdin = strings.NewReader(data)
	cmd.Stdout = f
	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
