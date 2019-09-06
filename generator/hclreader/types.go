package hclreader

// Config has a slice of rule that defines the analyzer.
type Config struct {
	Rules []Rule `hcl:"rule,block"`
}

// Rule has information to generate analyzer.
type Rule struct {
	Name    string  `hcl:"name,label"`
	Package string  `hcl:"package"`
	Doc     string  `hcl:"doc"`
	Message *string `hcl:"message,optional"`
	Types   []Type  `hcl:"type,block"`
	Funcs   []Func  `hcl:"func,block"`
}

// Type is a block to check the value of type is properly used.
type Type struct {
	Name    string   `hcl:"name,label"`
	Shoulds []string `hcl:"should"`
}

// Func is a block to check the function including method is properly used.
type Func struct {
	Name     string   `hcl:"name,label"`
	Receiver *string  `hcl:"receiver,optional"`
	Befores  []Before `hcl:"before,block"`
	Afters   []After  `hcl:"after,block"`
}

// Before is a block to check the specified function is called before the function of parent block is called.
type Before struct {
	Name   string   `hcl:"func,label"`
	Return []string `hcl:"return,optional"`
}

// After is a block to check the specified function is called after the function of parent block is called.
type After struct {
	Name string `hcl:"func,label"`
}
