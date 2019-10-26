package dsl_test

import (
	"fmt"
	. "github.com/pingcap/check"
	. "github.com/pingcap/parser/dsl"
	"github.com/pingcap/tidb/util/testleak"
)

var _ = Suite(&testDslSuite{})

type testDslSuite struct{}

func (s *testDslSuite) TestDSLParser(c *C) {
	defer testleak.AfterTest(c)()

	result, err := ExprParser(`lfkdsk: 111 and bdsl: [1 TO 1000] or abc: "lllll"`)
	if err == nil {
		fmt.Print(err)
	}

	c.Assert(err, IsNil)
	c.Assert(result, IsNil)
}

func (s *testDslSuite) TestDSLParser1(c *C) {
	defer testleak.AfterTest(c)()

	result, err := ExprParser(`lfkdsk: 111 and bdsl: [1 TO 1000] or abc: "lllll"`)
	if err == nil {
		fmt.Print(err)
	}

	c.Assert(err, IsNil)
	c.Assert(result, IsNil)
}
