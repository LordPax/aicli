package lang

import "github.com/LordPax/aicli/utils"

var EN_STRINGS = LangString{
	"usage":                   "CLI toot to use ai model",
	"output-desc":             "Output directory",
	"output-dir-empty":        "Output directory is empty",
	"silent":                  "Disable printing log to stdout",
	"no-args":                 "No arguments provided",
	"no-command":              "No command provided",
	"unknown-sdk":             "Unknown sdk \"%s\"",
	"sdk-model-usage":         "Select a model",
	"text-usage":              "Generate text from a prompt",
	"text-sdk-usage":          "Select a sdk",
	"text-temp-usage":         "Set temperature",
	"text-system-usage":       "Instruction with role system (use \"-\" for stdin)",
	"text-history-usage":      "Select a history",
	"text-clear-usage":        "Clear history",
	"text-file-usage":         "Text file to use",
	"text-input":              "(\"exit\" to quit) " + utils.Blue + "user> " + utils.Reset,
	"text-list-history-usage": "List history",
	"type-required":           "Type is required",
	"apiKey-required":         "API key is required",
	"empty-file":              "File \"%s\" is empty",
	"empty-history":           "History \"%s\" is empty\n",
}
