package models

type Options struct {
	ProjectPath string `short:"p" long:"project" description:"Path to unity project" default:"."`
	Output      string `short:"o" long:"output" description:"Output file"`
	RaiseError  bool   `short:"e" long:"raise-error" description:"Exit as error if even one is listed"`
}
