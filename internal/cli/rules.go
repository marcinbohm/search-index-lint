package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/marcinbohm/search-index-preflight/internal/diffrules"
	"github.com/marcinbohm/search-index-preflight/internal/rules"
)

func runRules(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 || isHelp(args[0]) {
		writeRulesHelp(stdout)
		return exitSuccess
	}

	switch args[0] {
	case "list":
		return runRulesList(args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "unknown rules command %q\n\n", args[0])
		writeRulesHelp(stderr)
		return exitUsage
	}
}

func runRulesList(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("rules list", flag.ContinueOnError)
	flags.SetOutput(stderr)

	var format string
	var family string
	flags.StringVar(&format, "format", "console", "Output format: console or json")
	flags.StringVar(&family, "family", "all", "Rule family: all, lint, or diff")
	flags.Usage = func() { writeRulesListHelp(flags.Output()) }

	if err := flags.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return exitSuccess
		}
		return exitUsage
	}
	if flags.NArg() != 0 {
		fmt.Fprintln(stderr, "rules list does not accept positional arguments")
		return exitUsage
	}
	if format != "console" && format != "json" {
		fmt.Fprintf(stderr, "invalid --format %q; expected console or json\n", format)
		return exitUsage
	}
	if family != "all" && family != "lint" && family != "diff" {
		fmt.Fprintf(stderr, "invalid --family %q; expected all, lint, or diff\n", family)
		return exitUsage
	}

	items, err := collectRuleListItems(family)
	if err != nil {
		fmt.Fprintf(stderr, "list rules: %v\n", err)
		return exitInternal
	}

	if format == "json" {
		output := struct {
			Rules []ruleListItem `json:"rules"`
		}{Rules: items}
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(output); err != nil {
			fmt.Fprintf(stderr, "write rules JSON: %v\n", err)
			return exitInternal
		}
		return exitSuccess
	}

	writeRuleListConsole(stdout, items)
	return exitSuccess
}

type ruleListItem struct {
	ID          string `json:"id"`
	Family      string `json:"family"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Severity    string `json:"severity"`
	Confidence  string `json:"confidence"`
	Determinism string `json:"determinism"`
	Description string `json:"description"`
}

func collectRuleListItems(family string) ([]ruleListItem, error) {
	var items []ruleListItem
	if family == "all" || family == "lint" {
		registry, err := rules.BuiltinRegistry()
		if err != nil {
			return nil, err
		}
		for _, rule := range registry.List() {
			metadata := rule.Metadata()
			items = append(items, ruleListItem{
				ID:          metadata.ID,
				Family:      "lint",
				Name:        metadata.Name,
				Category:    metadata.Category,
				Severity:    string(metadata.Severity),
				Confidence:  string(metadata.Confidence),
				Determinism: string(metadata.Determinism),
				Description: string(metadata.Description),
			})
		}
	}
	if family == "all" || family == "diff" {
		registry, err := diffrules.BuiltinRegistry()
		if err != nil {
			return nil, err
		}
		for _, rule := range registry.List() {
			metadata := rule.Metadata()
			items = append(items, ruleListItem{
				ID:          metadata.ID,
				Family:      "diff",
				Name:        metadata.Name,
				Category:    metadata.Category,
				Severity:    string(metadata.Severity),
				Confidence:  string(metadata.Confidence),
				Determinism: string(metadata.Determinism),
				Description: string(metadata.Description),
			})
		}
	}
	return items, nil
}

func writeRuleListConsole(w io.Writer, items []ruleListItem) {
	table := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintln(table, "ID\tFAMILY\tSEVERITY\tCATEGORY\tNAME")
	for _, item := range items {
		fmt.Fprintf(table, "%s\t%s\t%s\t%s\t%s\n", item.ID, item.Family, item.Severity, item.Category, item.Name)
	}
	table.Flush()
}

func writeRulesHelp(w io.Writer) {
	fmt.Fprint(w, `Usage:
  search-index-preflight rules <command>

Available Commands:
  list        List available rules
`)
}

func writeRulesListHelp(w io.Writer) {
	fmt.Fprint(w, `Usage:
  search-index-preflight rules list [flags]

List public lint and diff rule metadata.

Flags:
  --format <format>   Output format: console or json
  --family <family>   Rule family: all, lint, or diff
`)
}
