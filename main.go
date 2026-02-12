//	For API variables see: API.go

package main

import (
	"os"
	"fmt"
	"context"
	"encoding/json"

	"github.com/anthropics/anthropic-sdk-go"
)



func main() {
	jsonFile, err := os.Open(./JSONS/MOCK_DATA.json)
	if err != nil {
		fmt.println("err: a file failed to open")
		return
	}
	
//	//	or alternitvely for a lardge DB:
//	dir := "./JSONS/"
//	index, err := os.ReadDir(dir)
//	if err != nil {
//		panic(err)
//	}
//
//	for i := 0; i < index; i++ {
//
//		maximumJsonDigits := 0
//
//		if index == 0 {
//			fmt.Println("err: No JSON found in DB")
//			return
//		} else if {
//			maximumJsonDigits = len(strconv.Itoa(index)) //uses strconv library
//		} else {
//			fmt.Println("errL unknown")
//			return
//		}
//
//		jsonFile := fmt.Sprintf("assets/MOCK_DATA%0*d.png", maximumJsonDigits, i)
//	}
//	//	Then you can use all the code that runs in the file below and call jsonFile
//	//	in instead of the json above.


//	todo: Figure out if you are passing info to file properly & how to write to a json
//	file based of the input of claude.
fileUploadResult, err := client.Beta.Files.Upload(ctx, anthropic.BetaFileUploadParams{
		File:  anthropic.File(jsonFile, "file.txt", "text/json"),
		Betas: []anthropic.AnthropicBeta{anthropic.AnthropicBetaFilesAPI2025_04_14},
	})

	if err != nil {
		fmt.Printf("Error uploading file: %v\n", err)
		return
	}
	
	content := "Write me a summary of my file.txt file in the style of a Shakespearean sonnet.\n\n"
	println("[user]: " + content)

	message, err := client.Beta.Messages.New(ctx, anthropic.BetaMessageNewParams{
		MaxTokens: 1024,
		Messages: []anthropic.BetaMessageParam{
			anthropic.NewBetaUserMessage(
				anthropic.NewBetaTextBlock(content),
				anthropic.NewBetaDocumentBlock(anthropic.BetaFileDocumentSourceParam{
					FileID: fileUploadResult.ID,
				}),
			),
		},
		
		Model: modelNameAPI,
		Betas: []anthropic.AnthropicBeta{anthropic.AnthropicBetaFilesAPI2025_04_14},
	})

	if err != nil {
		fmt.Printf("Err creating message: %v\n", err)
		return
	}

	println("[assistant]: " + message.Content[0].Text + message.StopSequence)
}
