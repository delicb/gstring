package gst

import (
	"fmt"
	"io"
)

var defaultFormat = []rune("%s")

// Receives format string in gstring format and map of arguments and translates
// them to format and arguments compatible with standard golang formatting.
func gformat(format string, args map[string]interface{}) (string, []interface{}) {
	// holder for new format string - capacity as length of provided string
	// should be enough not to resize during appending, since expected length
	// of new format is smaller then provided one (names are removed)
	var new_format = make([]rune, 0, len(format))

	// flag that indicates if current place in format string in inside { }
	var in_format = false

	// flag that indicates if current place is format string in inside { } and after :
	var in_args = false

	// temp slice for holding name in current format
	var current_name_runes = make([]rune, 0, 10)

	// temp slice for holding args in current format
	var current_args_runes = make([]rune, 0, 10)

	var new_format_params []interface{}

	for _, ch := range format {
		switch ch {
		case '{':
			// Not supporting escaping { for now
			if in_format {
				panic("Invalid format string!")
			}
			in_format = true
		case '}':
			// Not supporting escaping } for now
			if !in_format {
				panic("Invalid format string")
			}
			if in_format {
				if len(current_args_runes) > 0 {
					// append formatting argumenst to new_format directly
					new_format = append(new_format, current_args_runes...)
				} else {
					// if no arguments are supplied, use default ones
					new_format = append(new_format, defaultFormat...)
				}
				// reset format args for new iteration
				current_args_runes = current_args_runes[0:0]
			}

			var name string
			if len(current_name_runes) == 0 {
				name = "EMPTY_PLACEHOLDER"
			} else {
				name = string(current_name_runes)
			}
			// reset name runes for next interation
			current_name_runes = current_name_runes[0:0]

			// get value from provided args and append it to new_format_args
			val, ok := args[name]
			if !ok {
				val = fmt.Sprintf("%%MISSING=%s", name)
			}
			new_format_params = append(new_format_params, val)

			// reset flags
			in_format = false
			in_args = false
		case ':':
			if in_format {
				in_args = true
			}
		default:
			if in_format {
				if in_args {
					current_args_runes = append(current_args_runes, ch)
				} else {
					current_name_runes = append(current_name_runes, ch)
				}
			} else {
				new_format = append(new_format, ch)
			}
		}
	}
	return string(new_format), new_format_params
}

// interface similar to "fmt" package
// Errorm returns error instance with formatted error message.
// This is same as fmt.Errorf, but uses gstring formatting.
func Errorm(format string, args map[string]interface{}) error {
	f, a := gformat(format, args)
	return fmt.Errorf(f, a...)
}

// Fprintm writes formatted string to provided writer.
// This is same as fmt.Fprintf, but uses gstring formatting.
func Fprintm(w io.Writer, format string, args map[string]interface{}) (n int, err error) {
	f, a := gformat(format, args)
	return fmt.Fprintf(w, f, a...)
}

// Printf prints formatted string to stdout.
// This is same as fmt.Printf, but uses gstring formatting.
func Printm(format string, args map[string]interface{}) (n int, err error) {
	f, a := gformat(format, args)
	return fmt.Printf(f, a...)
}

// Sprintf returns formatted string.
// This is same as fmt.Sprintf, but uses gstring formatting.
func Sprintm(format string, args map[string]interface{}) string {
	f, a := gformat(format, args)
	return fmt.Sprintf(f, a...)
}
