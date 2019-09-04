package hclreader

// Config has a slice of rule that defines the analyzer
type Config []Rule

// Rule has information to generate analyzer
type Rule struct {
	Package *string `hcl:"package"`
	Name    string `hcl:"name"`
	Message *string `hcl:"message"`
	Doc     *string `hcl:"doc"`
}
