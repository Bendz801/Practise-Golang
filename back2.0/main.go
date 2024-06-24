package main  
  
import (  
    "fmt"  
    "net/http"  
    "os/exec"  
    "encoding/json"
    "github.com/gin-gonic/gin"  
    "bytes"
    "strings"
    "strconv"
)  

type TextRequestBody struct {
    Text string
}
  
func runPythonScript(str string) (string, error) {  
    cmd := exec.Command("python", "script.py", str)

    var out bytes.Buffer
    cmd.Stdout = &out
    var stderr bytes.Buffer
    cmd.Stderr = &stderr
    
    err := cmd.Run()
    if err != nil {
        fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
        return "", err
    }

    asciiCodesStr := strings.Split(out.String(), " ")
    temp := asciiCodesStr[:len(asciiCodesStr) - 2]
    asciiCode, err := stringSliceToIntSlice(temp)

    return asciiToString(asciiCode), nil 
}
  
func main() {  
    r := gin.Default()  
    
    r.Static("/client", "./client")  
    r.LoadHTMLGlob("templates/*") 
    
    r.GET("/", func(c *gin.Context) {  
        c.HTML(http.StatusOK, "index.html", gin.H{})  
    })  
     
    r.POST("/correction-proposal", func(c *gin.Context) {
        bodyByteArray, err := c.GetRawData()

        var body TextRequestBody
        json.Unmarshal(bodyByteArray, &body)
        
        output, err := runPythonScript(body.Text)

        if err != nil {  
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
            return  
        }
        fmt.Println(body)
        c.JSON(http.StatusOK, gin.H{"result": output})  
    })  
     
    r.Run(":8080")  
}

func asciiToString(asciiCodes []int) string {
    var s string
    for _, code := range asciiCodes {
        s += string(rune(code))
    }
    return s
}

func stringSliceToIntSlice(strSlice []string) ([]int, error) {
    intSlice := make([]int, 0, len(strSlice))
    for _, str := range strSlice {
      intValue, err := strconv.Atoi(str)
      if err != nil {
        return nil, fmt.Errorf("error converting string '%s' to int: %w", str, err)
      }
      intSlice = append(intSlice, intValue)
    }
    

    return intSlice, nil
  }