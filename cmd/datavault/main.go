// cmd/datavault/main.go
package main

import (
"flag"
"log"
"os"

"datavault/internal/datavault"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := datavault.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
