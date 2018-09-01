package main

import (
	"regexp"
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)

	g.Describe("generateID", func() {
		g.It("Should return a 20 char string", func() {
			g.Assert(len(generateID("go"))).Equal(20)
			g.Assert(len(generateID("python"))).Equal(20)
			g.Assert(len(generateID("javascript"))).Equal(20)
			g.Assert(len(generateID("pygolangscript"))).Equal(20)
		})

		g.It("Should follow $language-runner-$lowercasestring format", func() {
			formatCheck := func(str string) bool {
				matched, _ := regexp.MatchString("^[a-z]*-runner-.[a-z]*$", str)
				return matched
			}

			g.Assert(formatCheck(generateID("go"))).Equal(true)
			g.Assert(formatCheck(generateID("python"))).Equal(true)
			g.Assert(formatCheck(generateID("javascript"))).Equal(true)
		})
	})

}
