package main

type CodeGenerator struct {
	buffer string
}

func (generator *CodeGenerator) WriteString(text string) {
	generator.buffer += text
}
