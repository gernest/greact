package mad

import (
	"fmt"
)

func ExampleDescribe() {
	test := Describe("passing test", It("must pass", func(_ T) {
	}))
	test.Exec()
	suite := test.(*Suite).Result()
	fmt.Println(len(suite.FailedExpectations) == 0)
	// Output: true
}
