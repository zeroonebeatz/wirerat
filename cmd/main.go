package main

import (
	"fmt"
	"log"
	"strings"
	"wirerat/cmd/internal/dock"
	"wirerat/cmd/internal/sniffer"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	selectedContainer, err := dock.SelectContainer()
	if err != nil {
		return err
	}

	if selectedContainer.Interface == nil {
		return fmt.Errorf("container selection error")
	}

	fmt.Printf("Names: %s\n", strings.Join(selectedContainer.Names, ","))
	fmt.Printf("Gateway: %v\n", selectedContainer.Gateway)

	if err := sniffer.Sniff(selectedContainer.Interface); err != nil {
		return err
	}

	return nil
}

