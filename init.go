package main

import (
	"fmt"
	"os"
)

var cmdInit = &Command{
	Run:       runInit,
	UsageLine: "init [OPTIONS]",
	Short:     "Generate markdown template",
	Long: `
Options:
	-h, --help     Print usage
`,
}

func init() {
}

func runInit(args []string) int {

	if len(args) > 0 {
		fmt.Fprintln(os.Stderr, "Too many arguments given.")
		return 1
	}

	base := `# PROBLEM



## EXISTING ALTERNATIVES



# CUSTOMER SEGMENTS



## EARLY ADOPTERS



# UNIQUE VALUE PROPOSITION



## HIGH-LEVEL CONCEPT



# SOLUTION



# CHANNELS



# REVENUE STREAMS



# COST STRUCTURE



# KEY METRICS



# UNFAIR ADVANTAGE



`
	fmt.Printf(base)
	return 0
}
