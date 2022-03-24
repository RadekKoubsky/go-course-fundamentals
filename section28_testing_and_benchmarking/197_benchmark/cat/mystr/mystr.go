package mystr

import "strings"

func Cat(xs []string) string {
	/*
		The implementation of concatenation often involves off-by-one error.
		An off-by-one error or off-by-one bug (known by acronyms OBOE, OBO, OB1 and OBOB) is a logic error involving the discrete equivalent of a boundary condition.

		More specifically, it contains the special case of this error caled Fencepost error:
		A fencepost error is a specific type of off-by-one error.

		In this scenario, a fence with n sections will have n + 1 posts. Conversely, if the fence contains n posts, it will contain n − 1 sections.
		In our impl:
		"section" -> space separator
		"posts" -> words

		Number of spaces = len(posts) - 1
	*/
	s := xs[0]
	/*
		To avoid having separator at the end of the loop, first we need to add the first word (s := xs[0]), then loop over
		the rest of words (range xs[1:]) and add separator + word for each word
	*/
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

func Join(xs []string) string {
	return strings.Join(xs, " ")
}
