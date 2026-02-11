package cmd


import (
    "os"
    "fmt"

    "github.com/spf13/cobra"
)


type FormaterP func (string) ([]byte, error);

func stringToBinary(data string) ([]byte, error) {
	return []byte(data), nil
}


func intToBinary(data string) ([]byte, error) {
	return nil, nil
}


var ConvbCmd = &cobra.Command{
	Use:   "convb [data]",
	Short: "Convert Single Data to Binary format",
	Args:  cobra.MinimumNArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {
		formats := map[string]FormaterP{
			"string": stringToBinary,
			//"int": intToBinary,
		}

		// Get the format flag
		format, err := cmd.Flags().GetString("format");
		if err != nil {
			fmt.Printf("Err: %s\n", err);
			os.Exit(1);
		}

		// Get the formater based on format string
		formater, exists := formats[format];
		if exists == false {
			fmt.Printf("Err: Unsupported: %s\n", format);
			os.Exit(1);
		}

		// Convert to binary.
		data := args[0]
		bytes, err := formater(data)
		fmt.Printf("%s\n", bytes);
		for _, b := range bytes {
			fmt.Printf("bytes:%08b\n", b);
		}
	},
}


func init() {
	// FMI:
	// SIG: func (c *Command) Flags() *flag.FlagSet
	// SIG: func (f *FlagSet) Bool(name string, value bool, usage string) *bool
	// SIG: func (f *FlagSet) String(name string, value string, usage string) *string
	ConvbCmd.Flags().String("format", "string", "Input format ['string', 'int']");
    rootCmd.AddCommand(ConvbCmd)
}
