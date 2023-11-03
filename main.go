package main

import (
	"fmt"
	openai "go-comment/internal/connectors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	apiKey     string
	inputFile  string
	outputFile string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "go-comment",
		Short: "A CLI tool to add comments to your code",
		Run:   run,
	}

	rootCmd.PersistentFlags().StringVarP(&apiKey, "api-key", "k", "", "OpenAI API Key")
	rootCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", "Input file")
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "", "Output file")

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func run(cmd *cobra.Command, args []string) {
	code, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	client := openai.NewClient(apiKey)

	fileType := extractFileType(outputFile)

	prompt := getPrompt(fileType, string(code))

	generatingMessage := "Generating comments..."
	fmt.Printf("%s \n", generatingMessage)

	fmt.Println(prompt)

	generatedComments, err := client.GetCompletion(openai.CompletionParams{
		Model:     "gpt-3.5-turbo-16k-0613",
		Prompt:    prompt,
		MaxTokens: 200,
	})

	if err != nil {
		log.Fatalf("Failed to make the API call: %v", err)
	}

	extractedCode := extractCodeBetweenFences(generatedComments, fileType)
	fmt.Printf("extracted:  %v \n", extractedCode)
	writeCommentsToFile(extractedCode)

	fmt.Print("\n")
	fmt.Printf("Comments added to %s\n", outputFile)
}

func extractCodeBetweenFences(input, fileType string) string {
	separator := fmt.Sprintf("```%s", fileType)
	parts := strings.Split(input, separator)
	if len(parts) >= 2 {
		codeParts := strings.Split(parts[1], "```")
		if len(codeParts) >= 1 {
			return strings.TrimSpace(codeParts[0])
		}
	}
	return ""
}

func extractFileType(filename string) string {
	ext := filepath.Ext(filename)
	if ext != "" {
		// remove dot (.) from the extension
		return ext[1:]
	}
	return ""
}

func getPrompt(fileType, code string) string {
	return fmt.Sprintf("\n\nAdd descriptive comments to the code in this %s file, avoiding comments on package names and entry functions, follow the following format, the code should be between ``` : \n ```%s\n%s\n```\n", fileType, fileType, code)
}

func writeCommentsToFile(generatedComments string) {
	commentLines := strings.Split(generatedComments, "\n")

	outputFileHandle, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer outputFileHandle.Close()

	for _, line := range commentLines {
		_, err = outputFileHandle.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("Failed to write to the output file: %v", err)
		}
	}
}
