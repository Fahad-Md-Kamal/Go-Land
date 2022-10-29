package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"syscall/js"
)

func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func jsonWrapper() js.Func {
	jsonfunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			result := map[string]any{
				"error": "Invalid no of arguments passed",
			}
			return result
		}
		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			result := map[string]any{
				"error": "Unable to get document object",
			}
			return result
		}
		jsonOuputTextArea := jsDoc.Call("getElementById", "jsonoutput")
		if !jsonOuputTextArea.Truthy() {
			result := map[string]any{
				"error": "Unable to get output text area",
			}
			return result
		}
		inputJSON := args[0].String()
		pretty, err := prettyJson(inputJSON)
		if err != nil {
			errStr := fmt.Sprintf("unable to parse JSON. Error %s occurred\n", err)
			result := map[string]any{
				"error": errStr,
			}
			return result
		}
		jsonOuputTextArea.Set("value", pretty)
		return nil
	})
	return jsonfunc
}

func Calculate() js.Func {
	jsonfunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) < 2 {
			result := map[string]any{
				"error": "Invalid no of arguments passed",
			}
			return result
		}
		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			result := map[string]any{
				"error": "Unable to get document object",
			}
			return result
		}
		resultOutput := jsDoc.Call("getElementById", "calcoutput")
		if !resultOutput.Truthy() {
			errStr := fmt.Sprintf("Output field not found")
			result := map[string]any{
				"error": errStr,
			}
			return result
		}
		inpNum1, err := strconv.Atoi(args[0].String())
		if err != nil {
			errStr := fmt.Sprintf("unable to parse value. Error %s occurred\n", err)
			result := map[string]any{
				"error": errStr,
			}
			return result
		}
		inpNum2, err := strconv.Atoi(args[1].String())
		if err != nil {
			errStr := fmt.Sprintf("unable to parse value. Error %s occurred\n", err)
			result := map[string]any{
				"error": errStr,
			}
			return result
		}

		outputFmt := fmt.Sprintf("The sum of %d + %d = %d", inpNum1, inpNum2, inpNum1+inpNum2)
		resultOutput.Set("innerText", outputFmt)
		return nil
	})
	return jsonfunc
}

func main() {
	println("WASM Go Initialized")
	channel := make(chan struct{}, 0)
	js.Global().Set("formatJSON", jsonWrapper())
	js.Global().Set("calculate", Calculate())
	<-channel
}
