package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/crazy3lf/colorconv"
	"github.com/fatih/color"
)

/*
	CL - utility used to highlight words in files opened with it.
*/

type WordsArray []string

func (i *WordsArray) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *WordsArray) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var Words WordsArray
var Color string

func GetColor(c string) string {
	switch Color {
	case "red":
		return color.New(color.FgRed).SprintFunc()(c)
	default:
		Color = strings.Trim(Color, "#")
		r, g, b, err := colorconv.HexToRGB(Color)
		if err != nil {
			return c
		}
		return color.RGB(int(r), int(g), int(b)).SprintFunc()(c)
	}
}

func main() {
	flag.Var(&Words, "word", "Add words to be highlighted.")
	flag.StringVar(&Color, "color", "red", "Set color for highlight.")
	flag.Parse()
	args := flag.Args()

	var f *os.File
	var err error
	if len(args) == 0 {
		f = os.Stdin
	} else {
		f, err = os.OpenFile(os.Args[len(os.Args)-1], os.O_APPEND, os.ModePerm)
		if err != nil {
			os.Exit(-1)
		}
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for i, word := range words {
			if slices.Contains(Words, word) {
				words[i] = GetColor(word)
			}
		}
		fmt.Println(strings.Join(words, " "))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
