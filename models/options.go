package models

type Options struct {
	ProjectPath string `short:"p" long:"project" description:"Path to unity project" required:"yes"`
	Output      string `short:"o" long:"output" description:"Output file"`
}
