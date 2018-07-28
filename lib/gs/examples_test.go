package gs_test

import (
	"fmt"

	"github.com/gernest/vected/lib/gs"
)

func ExampleS_selector() {
	css := gs.S("a",
		gs.P("color", "green"),
	)
	fmt.Println(gs.ToString(css))
	//Output:
	// a {
	//   color:green;
	// }
}

func ExampleS_nested() {
	css := gs.S("a",
		gs.P("color", "blue"),
		gs.S("&:hover",
			gs.P("color", "blue"),
		),
	)
	fmt.Println(gs.ToString(css))
	// Output:
	// a {
	//   color:blue;
	// }
	// a:hover {
	//   color:blue;
	// }
}

func ExampleS_parent() {
	gs.S("root",
		gs.S(" & > child_1",
			gs.S("& > child2",
				gs.P("key", "value"),
			),
		),
	)
	css := gs.S("root",
		gs.S(" & > child_1",
			gs.S("& > child2",
				gs.P("key", "value"),
			),
		),
	)
	fmt.Println(gs.ToString(css))
	//Output:
}
func ExampleP_attribute() {
	css := gs.P("color", "green")
	fmt.Println(gs.ToString(css))
	// Output:
	// color:green;
}
