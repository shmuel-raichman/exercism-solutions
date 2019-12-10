/*
Package twofer simple package that get name and returns massage with that name.
*/
// B"H
package twofer

// import fmt package to format the massege.
import "fmt"

// ShareWith return massege with the givan name if name is mot empty or "you" if name is empty.
func ShareWith(name string) string {
	// assign name the value "you" if string is empty.
	if name == "" {
		name = "you"
	}
	// concat the massage
	return fmt.Sprintf("One for %s, one for me.", name)
	
}
