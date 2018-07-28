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
	css := gs.S("root",
		gs.S(" & > child_1",
			gs.S("& > child2",
				gs.P("key", "value"),
			),
		),
	)
	fmt.Println(gs.ToString(css))
	//Output:
	// root > child_1 > child2 {
	//   key:value;
	// }
}
func ExampleP_attribute() {
	css := gs.P("color", "green")
	fmt.Println(gs.ToString(css))
	// Output:
	// color:green;
}
