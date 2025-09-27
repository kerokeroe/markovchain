package main

import (
	"bufio" //to read the data
	"flag"
	"fmt" //to print out
	"io"
	"os" //to open the file
	"strings"
)

type App struct {
	PrefixLen      int
	WordsNumber    int
	StartingPrefix string
	SourceFile     string
}

func NewApp() (*App, error) {
	fs := flag.NewFlagSet("markovchain", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	prefixLen := fs.Int("l", 2, "length of the prefix")
	wordsNum := fs.Int("w", 10, "how many words to generate")
	startPrefix := fs.String("p", "", "what word(s) should it start from")
	help := fs.Bool("help", false, "do you need any help?")

	if err := fs.Parse(os.Args[1:]); err != nil {
		printUsage()
		return nil, fmt.Errorf("invalid flags: %w", err)
	}
	if *help {
		printUsage()
		os.Exit(0)
	}
	var filepath string
	if fs.NArg() > 0 {
		filepath = fs.Arg(0)
	}

	if *prefixLen < 1 || *prefixLen > 5 {
		return nil, fmt.Errorf("prefix length should be between 1 and 5")
	}
	if *wordsNum < 1 || *wordsNum > 10000 {
		return nil, fmt.Errorf("the number of generated words should be between 1 and 10000")
	}
	return &App{
		PrefixLen:      *prefixLen,
		WordsNumber:    *wordsNum,
		StartingPrefix: *startPrefix,
		SourceFile:     filepath,
	}, nil

}

func printUsage() {
	fmt.Println("Markov Chain text generator.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  markovchain [-w <N>] [-p <S>] [-l <N>] [file]")
	fmt.Println("  markovchain --help")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --help  Show this screen.")
	fmt.Println("  -w N    Number of maximum words")
	fmt.Println("  -p S    Starting prefix")
	fmt.Println("  -l N    Prefix length")
}

func (a *App) Run() error {
	var reader io.Reader
	if a.SourceFile != "" {
		file, err := os.Open(a.SourceFile)
		if err != nil {
			return fmt.Errorf("couldnt open the file %s: %w", a.SourceFile, err)
		}
		defer file.Close()
		reader = file
	} else {
		stat, err := os.Stdin.Stat()
		if err != nil {
			return fmt.Errorf("couldnt read from stdin: %w", err)
		}
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			reader = os.Stdin
		} else {
			return fmt.Errorf("no input bruh")
		}
	}

	//reading the file
	words, err := readWords(reader)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}
	if len(words) == 0 {
		return fmt.Errorf("input text is empty")
	}
	if a.PrefixLen > 0 && len(words) < a.PrefixLen {
		return fmt.Errorf("input text is too short")
	}

	//build chain
	builder, err := NewBuilder(a.PrefixLen)
	if err!=nil{
		return fmt.Errorf("smth wrong with prefix, couldnt create the builder: %w", err)
	}
	chain := 

}

type Builder struct{
	prefixLen int
}

//this function validates prefixLen and returns the builder
func NewBuilder (prefixL int) (*Builder, error){
	if prefixL < 0 {
		return nil, fmt.Errorf("prefix length cant be negative")
	}
	if prefixL >5 {
		return nil, fmt.Errorf("prefix should be shorter than 5")
	}
	return &Builder{prefixLen: prefixL}, nil
}

type Chain struct{
	Chain map[string][]string
	prefixLen int
}

//function that builds the chain
func (b Builder) Build (words []string) (*Chain, error){
	chainmap := make([string] string, len(words))

	if b.prefixLen == 0 {
		chainmap[""] = append(chainmap[""], words ...)
		return &Chain{Chain: chainmap, prefixLen: b.prefixLen}, nil
	}
	if b.prefixLen > len(words){
		return &Chain{Chain: chainmap, prefixLen: b.prefixLen}, nil
	}

	for i:=0; i<len(words)-b.prefixLen-1; i++{
		prefixSlice := words[i:i+b.prefixLen]
		key := strings.ToLower(strings.Join(prefixSlice, " "))
		suffix := words[i+b.prefixLen]
		chainmap[key] = append(chainmap[key], suffix)
	}
	return &Chain{Chain: chainmap, prefixLen: : b.prefixLen}, nil

}
func readWords(source io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(source)
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		//Catches tokenizer errors (including the “token too long” case if you hit the 64 KB limit).
		return nil, fmt.Errorf("scanning input faile: %w", err)
	}
	return words, nil
}

func writeWords() {
	//blabla
}
func main() {
	file, f_err := os.Open("example.txt")
	if f_err != nil {
		panic(f_err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		items := strings.Split(line, " ")
		for _, item := range items {
			fmt.Println(item)
		}
	}

	// if len(os.Args)<2{
	// 	fmt.Print("Error: no input text")
	// 	os.Exit(1)
	// }
}

//plan:
//1. flag reading (flags: -w, -p, -l, -help)
//2. program takes a text file -> create readfile function
//3. program can also take input from stdin
//4.
