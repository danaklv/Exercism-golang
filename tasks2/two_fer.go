package tasks2

import "fmt"


// ShareWith function return string with message, which you will say when you get extra cookie,
//if you know person's name, it will return "One for <name>, one for me," if you dont know the person's nemae, it will return "One for you, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)

}

