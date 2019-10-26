package dsl

import (
	"fmt"
	. "github.com/vektah/goparsify"
)

type ParserDSLNode interface {
	GetResult() string
}

type Between struct {
	ParserDSLNode
	From   string
	To     string
	Result string
}

func (b *Between) GetResult() string {
	return b.Result
}

type Term struct {
	ParserDSLNode
	Label  string
	Result string
}

func (t *Term) GetResult() string {
	return t.Result
}

type BinaryOp struct {
	ParserDSLNode
	Op string
	L  ParserDSLNode
	R  ParserDSLNode
}

func (b *BinaryOp) GetResult() string {
	return ""
}

var (
	_string = StringLit(`'`).Map(func(n *Result) {
		n.Result = n.Token
	})
	_number = NumberLit().Map(func(n *Result) {
		n.Result = fmt.Sprintf("%v", n.Result)
	})
	_between = Seq("[", _number, "TO", _number, "]").Map(func(n *Result) {
		ret := Between{}
		ret.From = n.Child[1].Result.(string)
		ret.To = n.Child[3].Result.(string)
		ret.Result = "between " + n.Child[1].Result.(string) + " and " + n.Child[3].Result.(string)
		n.Result = ret
	})
	_label  = Regex("[a-zA-Z][a-zA-Z0-9]*")
	_result = Any(_string, _number, _between)
	_term   = Seq(_label, ":", _result).Map(func(n *Result) {
		ret := Term{}
		ret.Label = n.Child[0].Token
		switch n.Child[2].Result.(type) {
		case string:
			ret.Result = ret.Label + " = " + n.Child[2].Result.(string)
		case Between:
			ret.Result = ret.Label + " " + n.Child[2].Result.(Between).Result
		}
		n.Result = ret
	})

	_op = Any(Bind("and", nil), Bind("or", nil))

	_andOp = Seq(&_term, Some(Seq(_op, &_term))).Map(func(n *Result) {
		result := n.Child[0].Result.(Term).Result
		others := n.Child[1].Child
		for _, sub := range others {
			result += " " + sub.Child[0].Token + " " + sub.Child[1].Result.(Term).Result
		}
		n.Result = result
	})

	DslPhaseParser = Maybe(_andOp)
)

func ExprParser(dslString string) (string, error) {
	result, err := Run(DslPhaseParser, dslString)
	if err != nil {
		return "error", err
	}

	return result.(string), nil
}
