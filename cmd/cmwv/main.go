package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

var (
	newlineSpaceNewlineRegexp = regexp.MustCompile(`. +\n|\t`)
	tabRegexp                 = regexp.MustCompile(`\t`)
)

func run() error {
	content, helpRequested, err := getInput(os.Args)
	if err != nil {
		return err
	}
	if helpRequested {
		fmt.Printf(`cmwv - %sonfig%sap %shitespace %salidator
cmwv highlights parts of the provided Configmap that cause it to render in a "raw" format. It's usually a problem when you want to
get it in yaml format, for example using kubectl: 'kubectl get configmap -n $NAMESPACE $CONFIG_MAP_NAME -o yaml'.
There are 2 main reasons for this behaviour: tab (\t) characters and the trailing whitespace. cmwv changes \t to \\t and " " into "_", both highlighted in red, to mark places which need to be adjusted in the source. It return with exit code 1 if it catches any problem.
`, color.HiRedString("C"), color.HiRedString("M"), color.HiRedString("w"), color.HiRedString("v"))
		return nil
	}

	configMap := &corev1.ConfigMap{}
	if err := yaml.Unmarshal(content, configMap); err != nil {
		return fmt.Errorf("while unmarshaling input content into ConfigMap struct: %s", err)
	}
	if len(configMap.Data) == 0 {
		// nothing to validate
		return nil
	}

	keysWithInvalidValues := []string{}
	for key, value := range configMap.Data {
		// fmt.Println(strings.NewReplacer(" ", ".").Replace(value))
		val, foundInvalidParts := highlightInvalidParts(value)
		if foundInvalidParts {
			color.HiGreen("key: %s", key)
			fmt.Println(strings.TrimSuffix(val, "\n"))
			keysWithInvalidValues = append(keysWithInvalidValues, key)
		}
	}

	switch len(keysWithInvalidValues) {
	case 0:
		return nil
	case 1:
		return fmt.Errorf(color.RedString("data key %s contains characters that make configMap unreadable, they are highlighted in red color. For more info run `cmwv help`", keysWithInvalidValues[0]))
	default:
		return fmt.Errorf(color.RedString("data keys %s contain characters that make configMap unreadable, they are highlighted in red color. For more info run `cmwv help`", strings.Join(keysWithInvalidValues, ",")))
	}
}

func highlightInvalidParts(input string) (string, bool) {
	if !newlineSpaceNewlineRegexp.MatchString(input) {
		return input, false
	}

	return newlineSpaceNewlineRegexp.ReplaceAllStringFunc(input, func(arg string) string {
		return strings.
			NewReplacer(
				" ", color.RedString("_"),
				"\t", color.RedString("\\t"),
			).
			Replace(arg)
	}), true
}

func getInput(osArgs []string) ([]byte, bool, error) {
	switch l := len(osArgs); l {
	case 1:
		return nil, false, fmt.Errorf("Please specify path to file with configMap as a first argument, or a \"-\" to read configMap from stdin")
	case 2: // nothing
	default:
		return nil, false, fmt.Errorf(`Provided %d arguments, please provide only 1, either a path to a file with Kubernetes configMap or "-" for stdin`, l-1)
	}

	inputArg := os.Args[1]
	switch inputArg {
	case "help", "-h", "--help":
		return nil, true, nil
	case "-":
		content, err := io.ReadAll(os.Stdin)
		return content, false, err
	default:
		content, err := os.ReadFile(inputArg)
		return content, false, err
	}
}
