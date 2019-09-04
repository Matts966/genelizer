package hclreader

// Config has a slice of rule that defines the analyzer
type Config struct {
	Rules []Rule `hcl:"rule,block"`
}

// Rule has information to generate analyzer
type Rule struct {
	Name    string  `hcl:"name,label"`
	Package string  `hcl:"package"`
	Doc     string  `hcl:"doc"`
	Message *string `hcl:"message,optional"`
}
