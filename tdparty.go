package main

import (
	"fmt"
	"log"
	"os"
)
import "gopkg.in/xmlpath.v2"

// Token represents a true dungeon token
type Token struct {
	name      string
	rarity    string
	neededFor []string
	madeFrom  []string
	usableBy  []string
}

func extractRepeated(path string, context *xmlpath.Node) (result []string) {
	compiledPath := xmlpath.MustCompile(path)

	iter := compiledPath.Iter(context)

	var repeated []string
	for iter.Next() {
		fmt.Println("Found something")
		node := iter.Node()
		repeated = append(repeated, node.String())
	}

	return repeated
}

func extractString(path string, context *xmlpath.Node) (result string, ok bool) {
	compiledPath := xmlpath.MustCompile(path)

	return compiledPath.String(context)
}

func main() {
	file, err := os.Open("/Users/alek/td_crawl/tokendb.com/token/2-scepter-of-might/index.html")
	if err != nil {
		log.Fatal(err)
	}

	root, err := xmlpath.ParseHTML(file)
	if err != nil {
		log.Fatal(err)
	}

	rarity, _ := extractString("//div[@class='dir-tax' and contains(.,'Rarity')]/a", root)
	name, _ := extractString("//span[@class='dir-title']", root)
	neededFor := extractRepeated("//div[@class='dir-tax' and contains(.,'Ingredient For')]/a", root)
	usableBy := extractRepeated("//div[@class='dir-tax' and contains(.,'Usable By')]/a", root)

	fmt.Printf("%v\n", rarity)
	fmt.Printf("%v\n", name)
	fmt.Printf("NeededFor: %v\n", neededFor)

	token := Token{name: name, rarity: rarity, neededFor: neededFor, usableBy: usableBy}
	fmt.Printf("%v\n", token)
}
