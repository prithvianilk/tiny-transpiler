package main

type Generator struct {
	buffer string
}

func (generator *Generator) WriteString(text string) {
	generator.buffer += text
}
