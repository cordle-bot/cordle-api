package router

type Method int

// Enum defining which type of Method it is.
const (
	Undefined Method = iota // 0
	Get                     // 1
	Post                    // 2
	Patch                   // 3
	Delete                  // 4
)
