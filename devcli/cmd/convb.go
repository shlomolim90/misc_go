package cmd


import (
    "os"
    "fmt"
    "strconv"

    "github.com/spf13/cobra"
)


type FormaterP func (string) ([]byte, error);

func stringToBinary(data string) ([]byte, error) {
	return []byte(data), nil
}


func intToBinary(data string) ([]byte, error) {

    // SIG: func ParseInt(s string, base int, bitSize int) (i int64, err error)
    // if base is 0, the true base is implied by the string's prefix following
    // the sign 0b, 0x 0o, 0.
    int_data, err := strconv.ParseInt(data, 0, 64);
    if err != nil {
        return nil, err;
    }

    return uint64toBinary(uint64(int_data));
}


func uint64toBinary(data uint64) ([]byte, error) {
    // 8 bits * 8 = 64 bits
    buf := make([]byte, 8)
    for i := 0; i < 8; i++ {
        var currentByte byte = 0
        for j := 0; j < 8; j++ {
            currentByte <<= 1;
            // 63: uint64 에서 최상위 비트를 취득하기 위함.
            if (data & (1 << 63)) != 0 {
                currentByte |= 1;
            }
            data <<= 1;
        }
        buf[i] = currentByte;
    }
    return buf, nil
}


var ConvbCmd = &cobra.Command{
	Use:   "convb [data]",
	Short: "Convert Single Data to Binary format",
	Args:  cobra.MinimumNArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {
		formats := map[string]FormaterP{
			"string": stringToBinary,
			"int": intToBinary,
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
		for _, b := range bytes {
			fmt.Printf("byte:%08b\n", b);
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
