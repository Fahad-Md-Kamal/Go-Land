# GO-WASM

Create new project for GO

```sh
    go mod init github.com/fahad-md-kamal/wasm/json-prettyfiy
```

Create Go project's starting main module in ``main.go``

```go
    func main() {
        println("WASM Go Initialized")
        channel := make(chan struct{}, 0)
        js.Global().Set("formatJSON", jsonWrapper())
        <-channel
    }
```

Here, we are making a chan instance so that it stays open for longer duration.

Now create a function that takes the JavaScript's document object and does the megic on it.

```go
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
        jsonfunc := js.FuncOf(func(this js.Value, args []js.Value) any{
            if len(args) != 1 {
                result := map[string]any{
                    "error": "Invalid no of arguments passed",
                }
                return result
            }
            jsDoc := js.Global().Get("document") // Gets the document object from JS.
            if !jsDoc.Truthy() {
                result := map[string]any{
                    "error": "Unable to get document object",
                }
                return result
            }
            jsonOuputTextArea := jsDoc.Call("getElementById", "jsonoutput") // From document object get's the element by id
            if !jsonOuputTextArea.Truthy() {
                result := map[string]any{
                    "error": "Unable to get output text area",
                }
                return result
            }
            inputJSON := args[0].String()
            fmt.Printf("input %s\n", inputJSON)
            pretty, err := prettyJson(inputJSON)
            if err != nil {
                errStr := fmt.Sprintf("unable to parse JSON. Error %s occurred\n", err)
                result := map[string]any{
                    "error": errStr,
                }
                return result
            }
            jsonOuputTextArea.Set("value", pretty) //Sets the value for output textare
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
```

Glue wasm JSON from GO environment

```sh
    GOOS=js GOARCH=wasm go build -o <output-file-location> <go-main-file-path>
    #e.g: GOOS=js GOARCH=wasm go build -o ./static/main.wasm .
```

---

**Start GO server**
Create a main file in server diractory. As a different module and as a main package.

```go
    package main

    import (
        "flag"
        "log"
        "net/http"
    )

    var (
        listen = flag.String("listen", ":8080", "listen address")
        dir    = flag.String("dir", "./static", "directory to serve")
    )

    func main() {
        flag.Parse()
        log.Printf("listening on %q...", *listen)
        err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
        log.Fatalln(err)
    }

```

Run the server and let it running.  

---
***Build the solution***

```sh
    #GOOS=js GOARCH=wasm go build -o <built-wasm file output location> <go-main-file-location>
    GOOS=js GOARCH=wasm go build -o ./static/main.wasm .
```

---

## HTML

Create the html template which will serve our project.

```html
    <html>

    <head>
        <meta charset="utf-8" />
        <script src="wasm_exec.js"></script>
        <script>
            const go = new Go();
            WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
                go.run(result.instance);
            });
        </script>
    </head>

    <body>
        <textarea id="jsoninput" name="jsoninput" cols="80" rows="20"></textarea>
        <input id="button" type="submit" name="button" value="pretty json" onclick="json(jsoninput.value)" />
        <textarea id="jsonoutput" name="jsonoutput" cols="80" rows="20"></textarea>
        <br>
        <br>
        <input id="num1" name="num1" type="text">
        <input id="num2" name="num2" type="text">
        <input id="button" type="submit" name="button" value="Calculate" onclick="calc()" />
        <p id="calcoutput"></p><br>
    </body>
    <script>
        var json = function (input) {
            var result = formatJSON(input);
            if ((result != null) && ('error' in result)) {
                console.log("Go return value", result)
                jsonoutput.value = ""
                alert(result.error)
            }
        }
        var calc = function (input) {
            num1= document.getElementById('num1').value
            num2= document.getElementById('num2').value
            var result = calculate(num1, num2);
            if ((result != null) && ('error' in result)) {
                console.log("Go return value", result)
                calcoutput.innerText = ""
                alert(result.error)
            }
        }
    </script>

    </html>

```

[Tutorial Site](https://golangbot.com/webassembly-using-go/)
