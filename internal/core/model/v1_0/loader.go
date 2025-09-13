package v1_0

import (
	"fmt"
	"io"
	"os"
)

// LoadFromFile loads ERPN from file, auto-detecting format
func LoadFromFile(filename string) (*ERPN, error) {
	format := DetectFormat(filename)
	if format == "" {
		return nil, fmt.Errorf("unsupported file format: %s", filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	return ParseFromReader(file, format)
}

// SaveToFile saves ERPN to file in specified format
func (e *ERPN) SaveToFile(filename, format string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer file.Close()

	return e.ToWriter(file, format)
}

// ConvertFile converts ERPN file from one format to another
func ConvertFile(inputFile, outputFile string) error {
	// Load from input file
	erpn, err := LoadFromFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to load input file: %w", err)
	}

	// Detect output format
	outputFormat := DetectFormat(outputFile)
	if outputFormat == "" {
		return fmt.Errorf("unsupported output format for file: %s", outputFile)
	}

	// Save to output file
	return erpn.SaveToFile(outputFile, outputFormat)
}

// ValidateERPN performs basic validation on ERPN structure
func ValidateERPN(erpn *ERPN) []error {
	var errors []error

	// Validate required fields
	if erpn.Type.Name == "" {
		errors = append(errors, fmt.Errorf("type.name is required"))
	}
	if erpn.Type.Version == "" {
		errors = append(errors, fmt.Errorf("type.version is required"))
	}

	// Validate semver pattern for versions
	if erpn.Type.Version != "" && !isValidSemver(erpn.Type.Version) {
		errors = append(errors, fmt.Errorf("type.version must be valid semver (major.minor.patch): %s", erpn.Type.Version))
	}

	// Validate content if present
	if erpn.Content != nil {
		if erpn.Content.Version != nil && !isValidSemver(*erpn.Content.Version) {
			errors = append(errors, fmt.Errorf("content.version must be valid semver: %s", *erpn.Content.Version))
		}

		// Validate timing impact percentages
		if erpn.Content.Timing != nil {
			if erpn.Content.Timing.ImpactPercent != nil {
				if *erpn.Content.Timing.ImpactPercent < 0 || *erpn.Content.Timing.ImpactPercent > 100 {
					errors = append(errors, fmt.Errorf("timing.impactPercent must be between 0 and 100"))
				}
			}

			// Validate authors impact percentages
			for i, author := range erpn.Content.Timing.Authors {
				if author.ImpactPercent != nil {
					if *author.ImpactPercent < 0 || *author.ImpactPercent > 100 {
						errors = append(errors, fmt.Errorf("author[%d].impactPercent must be between 0 and 100", i))
					}
				}
			}
		}
	}

	return errors
}

// isValidSemver checks if a string matches semantic versioning pattern
func isValidSemver(version string) bool {
	// Simple regex-like check for semver pattern: major.minor.patch
	parts := make([]string, 0, 3)
	current := ""

	for _, r := range version {
		if r == '.' {
			if current == "" {
				return false
			}
			parts = append(parts, current)
			current = ""
		} else if r >= '0' && r <= '9' {
			current += string(r)
		} else {
			return false
		}
	}

	if current != "" {
		parts = append(parts, current)
	}

	return len(parts) == 3
}

// PrintSummary prints a human-readable summary of the ERPN
func (e *ERPN) PrintSummary(w io.Writer) {
	fmt.Fprintf(w, "ERPN Summary:\n")
	fmt.Fprintf(w, "  Type: %s v%s\n", e.Type.Name, e.Type.Version)

	if e.Content != nil {
		if e.Content.Name != nil {
			fmt.Fprintf(w, "  Content: %s", *e.Content.Name)
			if e.Content.Version != nil {
				fmt.Fprintf(w, " v%s", *e.Content.Version)
			}
			fmt.Fprintf(w, "\n")
		}

		if e.Content.Essence != nil {
			if e.Content.Essence.Synthesis != nil && e.Content.Essence.Synthesis.Fundamental != nil {
				fmt.Fprintf(w, "  Synthesis: %s\n", *e.Content.Essence.Synthesis.Fundamental)
			}

			if e.Content.Essence.Contradiction != nil && e.Content.Essence.Contradiction.Fundamental != nil {
				if e.Content.Essence.Contradiction.Fundamental.Thesis != nil {
					fmt.Fprintf(w, "  Thesis: %s\n", *e.Content.Essence.Contradiction.Fundamental.Thesis)
				}
				if e.Content.Essence.Contradiction.Fundamental.Antithesis != nil {
					fmt.Fprintf(w, "  Antithesis: %s\n", *e.Content.Essence.Contradiction.Fundamental.Antithesis)
				}
			}
		}

		if e.Content.Timing != nil && len(e.Content.Timing.Authors) > 0 {
			fmt.Fprintf(w, "  Authors:\n")
			for _, author := range e.Content.Timing.Authors {
				if author.Name != nil {
					fmt.Fprintf(w, "    - %s", *author.Name)
					if author.Position != nil {
						fmt.Fprintf(w, " (%s)", *author.Position)
					}
					if author.ImpactPercent != nil {
						fmt.Fprintf(w, " - %.1f%% impact", *author.ImpactPercent)
					}
					fmt.Fprintf(w, "\n")
				}
			}
		}
	}

	// Show extra fields count
	extraCount := 0
	if len(e.Extra) > 0 {
		extraCount += len(e.Extra)
	}
	if len(e.Type.Extra) > 0 {
		extraCount += len(e.Type.Extra)
	}
	if e.Content != nil && len(e.Content.Extra) > 0 {
		extraCount += len(e.Content.Extra)
	}

	if extraCount > 0 {
		fmt.Fprintf(w, "  Extra fields: %d preserved\n", extraCount)
	}
}
