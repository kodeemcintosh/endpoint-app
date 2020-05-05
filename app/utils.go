package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

func ParseFlags() {
	flag.StringVar(&commandFile, "file", "", "Use this flag if you want to input a single command")

	flag.Parse()
}

func ReadCommandFile(fileLocation string) []string {
	var commandList []string

	readFile, err := os.Open(fileLocation)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		commandList = append(commandList, scanner.Text())
	}

	return commandList
}

func GenerateNodeList(p Path) NodeList {
	strNodes := strings.Split(string(p), "/")

	var nodes NodeList

	for i, name := range strNodes {
		n := Node{
			Name:  name,
			Depth: i + 1,
		}

		nodes = append(nodes, &n)
	}

	return nodes
}
