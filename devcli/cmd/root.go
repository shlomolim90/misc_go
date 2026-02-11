package cmd


import (
    "os"
    "fmt"

    "github.com/spf13/cobra"
)


// cobra command 객체 생성.
var rootCmd = &cobra.Command{
    Use:   "devcli",
    Short: "DevCLI is a sample command-line application",
    Long:  `DevCLI is a longer description of your CLI application.`,
}

// Execute
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
	fmt.Printf("Bye!\n");
}

// 본 cmd 패키지가 처음 로드될 때, 수행되는 함수.
// 본 cmd 패키지가 명령 (서브명령) 을 추가해야하는 경우 init() 함수를 사용.
func init() {
    //rootCmd.AddCommand(pask.PaskCmd)
}

