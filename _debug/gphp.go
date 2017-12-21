package main

import (
	"encoding/json"
	"fmt"
	"github.com/emilioastarita/gphp/ast"
	"github.com/emilioastarita/gphp/lexer"
	"github.com/emilioastarita/gphp/parser"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func printUsage() {
	fmt.Println("Usage " + os.Args[0] + " scan|parse|compare filename")
}

func main() {
	if len(os.Args) < 3 {
		printUsage()
		return
	}

	if os.Args[1] == "scan" {
		fnWalk(os.Args[2], printTokensFromFile)
	} else if os.Args[1] == "parse" {
		fnWalk(os.Args[2], printAstFromFile)
	} else if os.Args[1] == "compare" {
		fnWalk(os.Args[2], printDiffWithPhp)
	} else {
		printUsage()
	}
}

func printDiffWithPhp(filename string) {
	p := parser.Parser{}

	sourceCase, _ := ioutil.ReadFile(filename)
	sourceFile := p.ParseSourceFile(string(sourceCase), "")
	jsonSource, _ := json.Marshal(ast.Serialize(&sourceFile))
	resultCase := getMsParserOutput(filename)

	differ := diff.New()

	d, err := differ.Compare(jsonSource, resultCase)

	if err != nil {
		fmt.Println("Log comparing json sources:")
		fmt.Println(err)
		// ast.PrettyPrintJSON(jsonSource)
		fmt.Println("Test fail.")
		return
	}

	if d.Modified() {
		println("Fail: ", filename)
		var aJson map[string]interface{}
		json.Unmarshal(jsonSource, &aJson)

		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       false,
		}
		formatter := formatter.NewAsciiFormatter(aJson, config)
		diffString, _ := formatter.Format(d)
		fmt.Println("START DIFF")
		fmt.Println(diffString)
		fmt.Println("END DIFF")

	} else {
		println("Ok: ", filename)
	}
}

func getMsParserOutput(filename string) []byte {
	cmd := "/usr/bin/php"
	args := []string{"/home/emilio/pry/tolerant-php-parser/tools/debug.php", filename}
	out, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		log.Fatalln("Error exec: ", cmd, err, string(out))
	}
	return out
}

func printAstFromFile(filename string) {
	fmt.Println("AST of :", filename)
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Can't read file:", filename)
		panic(err)
	}
	content := string(data)
	printAst(content)
}

func printTokensFromFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Can't read file:", filename)
		panic(err)
	}
	content := string(data)

	stream := lexer.TokensStream{}
	stream.Source(content)
	stream.CreateTokens()
	stream.Debug()

}

func printAst(content string) {

	p := parser.Parser{}
	sourceFile := p.ParseSourceFile(content, "")

	jsonSource, err := json.Marshal(ast.Serialize(&sourceFile))

	if err != nil {
		println(err)
	} else {
		pretty, _ := ast.PrettyPrintJSON(jsonSource)
		fmt.Println(string((pretty)))
	}
}

func lexerWalk(filename string) {
	list, _ := filesOfDir(filename)
	println("Reading: ", len(list))
	for _, f := range list {
		scanTokens(f)
	}
}

func fnWalk(fileOrDir string, fn func(file string)) {
	fi, err := os.Stat(fileOrDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		list, _ := filesOfDir(fileOrDir)
		for _, file := range list {
			fn(file)
		}
	case mode.IsRegular():
		fn(fileOrDir)
	}
}

func filesOfDir(searchDir string) ([]string, error) {
	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && path[len(path)-4:] == ".php" {
			fileList = append(fileList, path)
		}
		return err
	})

	if e != nil {
		panic(e)
	}
	return fileList, nil
}

func scanTokens(file string) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Exit opening file", err)
		return
	}
	//fmt.Println("Data", string(dat));
	stream := lexer.TokensStream{}
	stream.Source(string(dat))
	stream.CreateTokens()

	fmt.Println("File: ", file, "Tokens: ", len(stream.Tokens))
	//for key, token := range tokens {
	//	fmt.Println(key, token)
	//}
}
