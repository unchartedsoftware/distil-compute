package result

// Code generated by peg -inline ./primitive/compute/result/complex_field.peg DO NOT EDIT.

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

const endSymbol rune = 1114112

/* The rule types inferred from the grammar are below. */
type pegRule uint8

const (
	ruleUnknown pegRule = iota
	ruleComplexField
	rulearray
	rulearray_contents
	ruletuple_contents
	ruleitem
	rulestring
	ruledquote_string
	rulesquote_string
	rulevalue
	rulesep
	rulews
	rulecomma
	rulelf
	rulecr
	ruleescdquote
	ruleescsquote
	rulesquote
	ruleobracket
	rulecbracket
	ruleoparen
	rulecparen
	rulenumber
	rulenegative
	ruledecimal_point
	ruletextdata
	ruleAction0
	ruleAction1
	ruleAction2
	ruleAction3
	rulePegText
	ruleAction4
	ruleAction5
	ruleAction6
)

var rul3s = [...]string{
	"Unknown",
	"ComplexField",
	"array",
	"array_contents",
	"tuple_contents",
	"item",
	"string",
	"dquote_string",
	"squote_string",
	"value",
	"sep",
	"ws",
	"comma",
	"lf",
	"cr",
	"escdquote",
	"escsquote",
	"squote",
	"obracket",
	"cbracket",
	"oparen",
	"cparen",
	"number",
	"negative",
	"decimal_point",
	"textdata",
	"Action0",
	"Action1",
	"Action2",
	"Action3",
	"PegText",
	"Action4",
	"Action5",
	"Action6",
}

type token32 struct {
	pegRule
	begin, end uint32
}

func (t *token32) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v", rul3s[t.pegRule], t.begin, t.end)
}

type node32 struct {
	token32
	up, next *node32
}

func (node *node32) print(w io.Writer, pretty bool, buffer string) {
	var print func(node *node32, depth int)
	print = func(node *node32, depth int) {
		for node != nil {
			for c := 0; c < depth; c++ {
				fmt.Fprintf(w, " ")
			}
			rule := rul3s[node.pegRule]
			quote := strconv.Quote(string(([]rune(buffer)[node.begin:node.end])))
			if !pretty {
				fmt.Fprintf(w, "%v %v\n", rule, quote)
			} else {
				fmt.Fprintf(w, "\x1B[34m%v\x1B[m %v\n", rule, quote)
			}
			if node.up != nil {
				print(node.up, depth+1)
			}
			node = node.next
		}
	}
	print(node, 0)
}

func (node *node32) Print(w io.Writer, buffer string) {
	node.print(w, false, buffer)
}

func (node *node32) PrettyPrint(w io.Writer, buffer string) {
	node.print(w, true, buffer)
}

type tokens32 struct {
	tree []token32
}

func (t *tokens32) Trim(length uint32) {
	t.tree = t.tree[:length]
}

func (t *tokens32) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens32) AST() *node32 {
	type element struct {
		node *node32
		down *element
	}
	tokens := t.Tokens()
	var stack *element
	for _, token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	if stack != nil {
		return stack.node
	}
	return nil
}

func (t *tokens32) PrintSyntaxTree(buffer string) {
	t.AST().Print(os.Stdout, buffer)
}

func (t *tokens32) WriteSyntaxTree(w io.Writer, buffer string) {
	t.AST().Print(w, buffer)
}

func (t *tokens32) PrettyPrintSyntaxTree(buffer string) {
	t.AST().PrettyPrint(os.Stdout, buffer)
}

func (t *tokens32) Add(rule pegRule, begin, end, index uint32) {
	tree, i := t.tree, int(index)
	if i >= len(tree) {
		t.tree = append(tree, token32{pegRule: rule, begin: begin, end: end})
		return
	}
	tree[i] = token32{pegRule: rule, begin: begin, end: end}
}

func (t *tokens32) Tokens() []token32 {
	return t.tree
}

type ComplexField struct {
	arrayElements

	Buffer string
	buffer []rune
	rules  [34]func() bool
	parse  func(rule ...int) error
	reset  func()
	Pretty bool
	tokens32
}

func (p *ComplexField) Parse(rule ...int) error {
	return p.parse(rule...)
}

func (p *ComplexField) Reset() {
	p.reset()
}

type textPosition struct {
	line, symbol int
}

type textPositionMap map[int]textPosition

func translatePositions(buffer []rune, positions []int) textPositionMap {
	length, translations, j, line, symbol := len(positions), make(textPositionMap, len(positions)), 0, 1, 0
	sort.Ints(positions)

search:
	for i, c := range buffer {
		if c == '\n' {
			line, symbol = line+1, 0
		} else {
			symbol++
		}
		if i == positions[j] {
			translations[positions[j]] = textPosition{line, symbol}
			for j++; j < length; j++ {
				if i != positions[j] {
					continue search
				}
			}
			break search
		}
	}

	return translations
}

type parseError struct {
	p   *ComplexField
	max token32
}

func (e *parseError) Error() string {
	tokens, err := []token32{e.max}, "\n"
	positions, p := make([]int, 2*len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p+1
		positions[p], p = int(token.end), p+1
	}
	translations := translatePositions(e.p.buffer, positions)
	format := "parse error near %v (line %v symbol %v - line %v symbol %v):\n%v\n"
	if e.p.Pretty {
		format = "parse error near \x1B[34m%v\x1B[m (line %v symbol %v - line %v symbol %v):\n%v\n"
	}
	for _, token := range tokens {
		begin, end := int(token.begin), int(token.end)
		err += fmt.Sprintf(format,
			rul3s[token.pegRule],
			translations[begin].line, translations[begin].symbol,
			translations[end].line, translations[end].symbol,
			strconv.Quote(string(e.p.buffer[begin:end])))
	}

	return err
}

func (p *ComplexField) PrintSyntaxTree() {
	if p.Pretty {
		p.tokens32.PrettyPrintSyntaxTree(p.Buffer)
	} else {
		p.tokens32.PrintSyntaxTree(p.Buffer)
	}
}

func (p *ComplexField) WriteSyntaxTree(w io.Writer) {
	p.tokens32.WriteSyntaxTree(w, p.Buffer)
}

func (p *ComplexField) Execute() {
	buffer, _buffer, text, begin, end := p.Buffer, p.buffer, "", 0, 0
	for _, token := range p.Tokens() {
		switch token.pegRule {

		case rulePegText:
			begin, end = int(token.begin), int(token.end)
			text = string(_buffer[begin:end])

		case ruleAction0:
			p.pushArray()
		case ruleAction1:
			p.popArray()
		case ruleAction2:
			p.pushArray()
		case ruleAction3:
			p.popArray()
		case ruleAction4:
			p.addElement(buffer[begin:end])
		case ruleAction5:
			p.addElement(buffer[begin:end])
		case ruleAction6:
			p.addElement(buffer[begin:end])

		}
	}
	_, _, _, _, _ = buffer, _buffer, text, begin, end
}

func Pretty(pretty bool) func(*ComplexField) error {
	return func(p *ComplexField) error {
		p.Pretty = pretty
		return nil
	}
}

func Size(size int) func(*ComplexField) error {
	return func(p *ComplexField) error {
		p.tokens32 = tokens32{tree: make([]token32, 0, size)}
		return nil
	}
}
func (p *ComplexField) Init(options ...func(*ComplexField) error) error {
	var (
		max                  token32
		position, tokenIndex uint32
		buffer               []rune
	)
	for _, option := range options {
		err := option(p)
		if err != nil {
			return err
		}
	}
	p.reset = func() {
		max = token32{}
		position, tokenIndex = 0, 0

		p.buffer = []rune(p.Buffer)
		if len(p.buffer) == 0 || p.buffer[len(p.buffer)-1] != endSymbol {
			p.buffer = append(p.buffer, endSymbol)
		}
		buffer = p.buffer
	}
	p.reset()

	_rules := p.rules
	tree := p.tokens32
	p.parse = func(rule ...int) error {
		r := 1
		if len(rule) > 0 {
			r = rule[0]
		}
		matches := p.rules[r]()
		p.tokens32 = tree
		if matches {
			p.Trim(tokenIndex)
			return nil
		}
		return &parseError{p, max}
	}

	add := func(rule pegRule, begin uint32) {
		tree.Add(rule, begin, position, tokenIndex)
		tokenIndex++
		if begin != position && position > max.end {
			max = token32{rule, begin, position}
		}
	}

	matchDot := func() bool {
		if buffer[position] != endSymbol {
			position++
			return true
		}
		return false
	}

	/*matchChar := func(c byte) bool {
		if buffer[position] == c {
			position++
			return true
		}
		return false
	}*/

	/*matchRange := func(lower byte, upper byte) bool {
		if c := buffer[position]; c >= lower && c <= upper {
			position++
			return true
		}
		return false
	}*/

	_rules = [...]func() bool{
		nil,
		/* 0 ComplexField <- <(array !.)> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				if !_rules[rulearray]() {
					goto l0
				}
				{
					position2, tokenIndex2 := position, tokenIndex
					if !matchDot() {
						goto l2
					}
					goto l0
				l2:
					position, tokenIndex = position2, tokenIndex2
				}
				add(ruleComplexField, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		/* 1 array <- <((ws* obracket Action0 ws* array_contents ws* cbracket Action1) / (ws* oparen Action2 ws* tuple_contents ws* cparen Action3))> */
		func() bool {
			position3, tokenIndex3 := position, tokenIndex
			{
				position4 := position
				{
					position5, tokenIndex5 := position, tokenIndex
				l7:
					{
						position8, tokenIndex8 := position, tokenIndex
						if !_rules[rulews]() {
							goto l8
						}
						goto l7
					l8:
						position, tokenIndex = position8, tokenIndex8
					}
					if !_rules[ruleobracket]() {
						goto l6
					}
					{
						add(ruleAction0, position)
					}
				l10:
					{
						position11, tokenIndex11 := position, tokenIndex
						if !_rules[rulews]() {
							goto l11
						}
						goto l10
					l11:
						position, tokenIndex = position11, tokenIndex11
					}
					{
						position12 := position
					l13:
						{
							position14, tokenIndex14 := position, tokenIndex
							if !_rules[rulews]() {
								goto l14
							}
							goto l13
						l14:
							position, tokenIndex = position14, tokenIndex14
						}
						if !_rules[ruleitem]() {
							goto l6
						}
					l15:
						{
							position16, tokenIndex16 := position, tokenIndex
							if !_rules[rulews]() {
								goto l16
							}
							goto l15
						l16:
							position, tokenIndex = position16, tokenIndex16
						}
					l17:
						{
							position18, tokenIndex18 := position, tokenIndex
							if !_rules[rulesep]() {
								goto l18
							}
						l19:
							{
								position20, tokenIndex20 := position, tokenIndex
								if !_rules[rulews]() {
									goto l20
								}
								goto l19
							l20:
								position, tokenIndex = position20, tokenIndex20
							}
							if !_rules[ruleitem]() {
								goto l18
							}
						l21:
							{
								position22, tokenIndex22 := position, tokenIndex
								if !_rules[rulews]() {
									goto l22
								}
								goto l21
							l22:
								position, tokenIndex = position22, tokenIndex22
							}
							goto l17
						l18:
							position, tokenIndex = position18, tokenIndex18
						}
						{
							position23, tokenIndex23 := position, tokenIndex
							if !_rules[rulesep]() {
								goto l23
							}
							goto l24
						l23:
							position, tokenIndex = position23, tokenIndex23
						}
					l24:
						add(rulearray_contents, position12)
					}
				l25:
					{
						position26, tokenIndex26 := position, tokenIndex
						if !_rules[rulews]() {
							goto l26
						}
						goto l25
					l26:
						position, tokenIndex = position26, tokenIndex26
					}
					if !_rules[rulecbracket]() {
						goto l6
					}
					{
						add(ruleAction1, position)
					}
					goto l5
				l6:
					position, tokenIndex = position5, tokenIndex5
				l28:
					{
						position29, tokenIndex29 := position, tokenIndex
						if !_rules[rulews]() {
							goto l29
						}
						goto l28
					l29:
						position, tokenIndex = position29, tokenIndex29
					}
					if !_rules[ruleoparen]() {
						goto l3
					}
					{
						add(ruleAction2, position)
					}
				l31:
					{
						position32, tokenIndex32 := position, tokenIndex
						if !_rules[rulews]() {
							goto l32
						}
						goto l31
					l32:
						position, tokenIndex = position32, tokenIndex32
					}
					{
						position33 := position
						if !_rules[ruleitem]() {
							goto l3
						}
					l36:
						{
							position37, tokenIndex37 := position, tokenIndex
							if !_rules[rulews]() {
								goto l37
							}
							goto l36
						l37:
							position, tokenIndex = position37, tokenIndex37
						}
						if !_rules[rulecomma]() {
							goto l3
						}
					l38:
						{
							position39, tokenIndex39 := position, tokenIndex
							if !_rules[rulews]() {
								goto l39
							}
							goto l38
						l39:
							position, tokenIndex = position39, tokenIndex39
						}
					l34:
						{
							position35, tokenIndex35 := position, tokenIndex
							if !_rules[ruleitem]() {
								goto l35
							}
						l40:
							{
								position41, tokenIndex41 := position, tokenIndex
								if !_rules[rulews]() {
									goto l41
								}
								goto l40
							l41:
								position, tokenIndex = position41, tokenIndex41
							}
							if !_rules[rulecomma]() {
								goto l35
							}
						l42:
							{
								position43, tokenIndex43 := position, tokenIndex
								if !_rules[rulews]() {
									goto l43
								}
								goto l42
							l43:
								position, tokenIndex = position43, tokenIndex43
							}
							goto l34
						l35:
							position, tokenIndex = position35, tokenIndex35
						}
						{
							position44, tokenIndex44 := position, tokenIndex
						l46:
							{
								position47, tokenIndex47 := position, tokenIndex
								if !_rules[rulews]() {
									goto l47
								}
								goto l46
							l47:
								position, tokenIndex = position47, tokenIndex47
							}
							if !_rules[ruleitem]() {
								goto l44
							}
						l48:
							{
								position49, tokenIndex49 := position, tokenIndex
								if !_rules[rulews]() {
									goto l49
								}
								goto l48
							l49:
								position, tokenIndex = position49, tokenIndex49
							}
							goto l45
						l44:
							position, tokenIndex = position44, tokenIndex44
						}
					l45:
						add(ruletuple_contents, position33)
					}
				l50:
					{
						position51, tokenIndex51 := position, tokenIndex
						if !_rules[rulews]() {
							goto l51
						}
						goto l50
					l51:
						position, tokenIndex = position51, tokenIndex51
					}
					if !_rules[rulecparen]() {
						goto l3
					}
					{
						add(ruleAction3, position)
					}
				}
			l5:
				add(rulearray, position4)
			}
			return true
		l3:
			position, tokenIndex = position3, tokenIndex3
			return false
		},
		/* 2 array_contents <- <(ws* item ws* (sep ws* item ws*)* sep?)> */
		nil,
		/* 3 tuple_contents <- <((item ws* comma ws*)+ (ws* item ws*)?)> */
		nil,
		/* 4 item <- <(array / string / (<value> Action4))> */
		func() bool {
			position55, tokenIndex55 := position, tokenIndex
			{
				position56 := position
				{
					position57, tokenIndex57 := position, tokenIndex
					if !_rules[rulearray]() {
						goto l58
					}
					goto l57
				l58:
					position, tokenIndex = position57, tokenIndex57
					{
						position60 := position
						{
							position61, tokenIndex61 := position, tokenIndex
							{
								position63 := position
								if !_rules[ruleescdquote]() {
									goto l62
								}
								{
									position64 := position
								l65:
									{
										position66, tokenIndex66 := position, tokenIndex
										{
											position67, tokenIndex67 := position, tokenIndex
											if !_rules[ruletextdata]() {
												goto l68
											}
											goto l67
										l68:
											position, tokenIndex = position67, tokenIndex67
											if !_rules[rulesquote]() {
												goto l69
											}
											goto l67
										l69:
											position, tokenIndex = position67, tokenIndex67
											if !_rules[rulelf]() {
												goto l70
											}
											goto l67
										l70:
											position, tokenIndex = position67, tokenIndex67
											if !_rules[rulecr]() {
												goto l71
											}
											goto l67
										l71:
											position, tokenIndex = position67, tokenIndex67
											if !_rules[ruleobracket]() {
												goto l72
											}
											goto l67
										l72:
											position, tokenIndex = position67, tokenIndex67
											if !_rules[rulecbracket]() {
												goto l73
											}
											goto l67
										l73:
											position, tokenIndex = position67, tokenIndex67
											if !_rules[ruleoparen]() {
												goto l74
											}
											goto l67
										l74:
											position, tokenIndex = position67, tokenIndex67
											if !_rules[rulecparen]() {
												goto l75
											}
											goto l67
										l75:
											position, tokenIndex = position67, tokenIndex67
											if !_rules[rulecomma]() {
												goto l66
											}
										}
									l67:
										goto l65
									l66:
										position, tokenIndex = position66, tokenIndex66
									}
									add(rulePegText, position64)
								}
								if !_rules[ruleescdquote]() {
									goto l62
								}
								{
									add(ruleAction5, position)
								}
								add(ruledquote_string, position63)
							}
							goto l61
						l62:
							position, tokenIndex = position61, tokenIndex61
							{
								position77 := position
								if !_rules[rulesquote]() {
									goto l59
								}
								{
									position78 := position
								l79:
									{
										position80, tokenIndex80 := position, tokenIndex
										{
											position81, tokenIndex81 := position, tokenIndex
											{
												position83 := position
												if buffer[position] != rune('\\') {
													goto l82
												}
												position++
												if buffer[position] != rune('\'') {
													goto l82
												}
												position++
												add(ruleescsquote, position83)
											}
											goto l81
										l82:
											position, tokenIndex = position81, tokenIndex81
											if !_rules[ruleescdquote]() {
												goto l84
											}
											goto l81
										l84:
											position, tokenIndex = position81, tokenIndex81
											if !_rules[ruletextdata]() {
												goto l85
											}
											goto l81
										l85:
											position, tokenIndex = position81, tokenIndex81
											if !_rules[rulelf]() {
												goto l86
											}
											goto l81
										l86:
											position, tokenIndex = position81, tokenIndex81
											if !_rules[rulecr]() {
												goto l87
											}
											goto l81
										l87:
											position, tokenIndex = position81, tokenIndex81
											if !_rules[ruleobracket]() {
												goto l88
											}
											goto l81
										l88:
											position, tokenIndex = position81, tokenIndex81
											if !_rules[rulecbracket]() {
												goto l89
											}
											goto l81
										l89:
											position, tokenIndex = position81, tokenIndex81
											if !_rules[ruleoparen]() {
												goto l90
											}
											goto l81
										l90:
											position, tokenIndex = position81, tokenIndex81
											if !_rules[rulecparen]() {
												goto l80
											}
										}
									l81:
										goto l79
									l80:
										position, tokenIndex = position80, tokenIndex80
									}
									add(rulePegText, position78)
								}
								if !_rules[rulesquote]() {
									goto l59
								}
								{
									add(ruleAction6, position)
								}
								add(rulesquote_string, position77)
							}
						}
					l61:
						add(rulestring, position60)
					}
					goto l57
				l59:
					position, tokenIndex = position57, tokenIndex57
					{
						position92 := position
						{
							position93 := position
							{
								position94, tokenIndex94 := position, tokenIndex
								if !_rules[rulenegative]() {
									goto l94
								}
								goto l95
							l94:
								position, tokenIndex = position94, tokenIndex94
							}
						l95:
							if !_rules[rulenumber]() {
								goto l55
							}
						l96:
							{
								position97, tokenIndex97 := position, tokenIndex
								if !_rules[rulenumber]() {
									goto l97
								}
								goto l96
							l97:
								position, tokenIndex = position97, tokenIndex97
							}
							{
								position98, tokenIndex98 := position, tokenIndex
								{
									position100 := position
									if buffer[position] != rune('.') {
										goto l98
									}
									position++
									add(ruledecimal_point, position100)
								}
								if !_rules[rulenumber]() {
									goto l98
								}
							l101:
								{
									position102, tokenIndex102 := position, tokenIndex
									if !_rules[rulenumber]() {
										goto l102
									}
									goto l101
								l102:
									position, tokenIndex = position102, tokenIndex102
								}
								{
									position103, tokenIndex103 := position, tokenIndex
									if !_rules[rulenegative]() {
										goto l103
									}
									goto l104
								l103:
									position, tokenIndex = position103, tokenIndex103
								}
							l104:
							l105:
								{
									position106, tokenIndex106 := position, tokenIndex
									if !_rules[rulenumber]() {
										goto l106
									}
									goto l105
								l106:
									position, tokenIndex = position106, tokenIndex106
								}
								goto l99
							l98:
								position, tokenIndex = position98, tokenIndex98
							}
						l99:
							add(rulevalue, position93)
						}
						add(rulePegText, position92)
					}
					{
						add(ruleAction4, position)
					}
				}
			l57:
				add(ruleitem, position56)
			}
			return true
		l55:
			position, tokenIndex = position55, tokenIndex55
			return false
		},
		/* 5 string <- <(dquote_string / squote_string)> */
		nil,
		/* 6 dquote_string <- <(escdquote <(textdata / squote / lf / cr / obracket / cbracket / oparen / cparen / comma)*> escdquote Action5)> */
		nil,
		/* 7 squote_string <- <(squote <(escsquote / escdquote / textdata / lf / cr / obracket / cbracket / oparen / cparen)*> squote Action6)> */
		nil,
		/* 8 value <- <(negative? number+ (decimal_point number+ negative? number*)?)> */
		nil,
		/* 9 sep <- <(comma / lf)> */
		func() bool {
			position112, tokenIndex112 := position, tokenIndex
			{
				position113 := position
				{
					position114, tokenIndex114 := position, tokenIndex
					if !_rules[rulecomma]() {
						goto l115
					}
					goto l114
				l115:
					position, tokenIndex = position114, tokenIndex114
					if !_rules[rulelf]() {
						goto l112
					}
				}
			l114:
				add(rulesep, position113)
			}
			return true
		l112:
			position, tokenIndex = position112, tokenIndex112
			return false
		},
		/* 10 ws <- <' '> */
		func() bool {
			position116, tokenIndex116 := position, tokenIndex
			{
				position117 := position
				if buffer[position] != rune(' ') {
					goto l116
				}
				position++
				add(rulews, position117)
			}
			return true
		l116:
			position, tokenIndex = position116, tokenIndex116
			return false
		},
		/* 11 comma <- <','> */
		func() bool {
			position118, tokenIndex118 := position, tokenIndex
			{
				position119 := position
				if buffer[position] != rune(',') {
					goto l118
				}
				position++
				add(rulecomma, position119)
			}
			return true
		l118:
			position, tokenIndex = position118, tokenIndex118
			return false
		},
		/* 12 lf <- <'\n'> */
		func() bool {
			position120, tokenIndex120 := position, tokenIndex
			{
				position121 := position
				if buffer[position] != rune('\n') {
					goto l120
				}
				position++
				add(rulelf, position121)
			}
			return true
		l120:
			position, tokenIndex = position120, tokenIndex120
			return false
		},
		/* 13 cr <- <'\r'> */
		func() bool {
			position122, tokenIndex122 := position, tokenIndex
			{
				position123 := position
				if buffer[position] != rune('\r') {
					goto l122
				}
				position++
				add(rulecr, position123)
			}
			return true
		l122:
			position, tokenIndex = position122, tokenIndex122
			return false
		},
		/* 14 escdquote <- <'"'> */
		func() bool {
			position124, tokenIndex124 := position, tokenIndex
			{
				position125 := position
				if buffer[position] != rune('"') {
					goto l124
				}
				position++
				add(ruleescdquote, position125)
			}
			return true
		l124:
			position, tokenIndex = position124, tokenIndex124
			return false
		},
		/* 15 escsquote <- <('\\' '\'')> */
		nil,
		/* 16 squote <- <'\''> */
		func() bool {
			position127, tokenIndex127 := position, tokenIndex
			{
				position128 := position
				if buffer[position] != rune('\'') {
					goto l127
				}
				position++
				add(rulesquote, position128)
			}
			return true
		l127:
			position, tokenIndex = position127, tokenIndex127
			return false
		},
		/* 17 obracket <- <'['> */
		func() bool {
			position129, tokenIndex129 := position, tokenIndex
			{
				position130 := position
				if buffer[position] != rune('[') {
					goto l129
				}
				position++
				add(ruleobracket, position130)
			}
			return true
		l129:
			position, tokenIndex = position129, tokenIndex129
			return false
		},
		/* 18 cbracket <- <']'> */
		func() bool {
			position131, tokenIndex131 := position, tokenIndex
			{
				position132 := position
				if buffer[position] != rune(']') {
					goto l131
				}
				position++
				add(rulecbracket, position132)
			}
			return true
		l131:
			position, tokenIndex = position131, tokenIndex131
			return false
		},
		/* 19 oparen <- <'('> */
		func() bool {
			position133, tokenIndex133 := position, tokenIndex
			{
				position134 := position
				if buffer[position] != rune('(') {
					goto l133
				}
				position++
				add(ruleoparen, position134)
			}
			return true
		l133:
			position, tokenIndex = position133, tokenIndex133
			return false
		},
		/* 20 cparen <- <')'> */
		func() bool {
			position135, tokenIndex135 := position, tokenIndex
			{
				position136 := position
				if buffer[position] != rune(')') {
					goto l135
				}
				position++
				add(rulecparen, position136)
			}
			return true
		l135:
			position, tokenIndex = position135, tokenIndex135
			return false
		},
		/* 21 number <- <([a-z] / [A-Z] / [0-9])> */
		func() bool {
			position137, tokenIndex137 := position, tokenIndex
			{
				position138 := position
				{
					position139, tokenIndex139 := position, tokenIndex
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l140
					}
					position++
					goto l139
				l140:
					position, tokenIndex = position139, tokenIndex139
					if c := buffer[position]; c < rune('A') || c > rune('Z') {
						goto l141
					}
					position++
					goto l139
				l141:
					position, tokenIndex = position139, tokenIndex139
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l137
					}
					position++
				}
			l139:
				add(rulenumber, position138)
			}
			return true
		l137:
			position, tokenIndex = position137, tokenIndex137
			return false
		},
		/* 22 negative <- <'-'> */
		func() bool {
			position142, tokenIndex142 := position, tokenIndex
			{
				position143 := position
				if buffer[position] != rune('-') {
					goto l142
				}
				position++
				add(rulenegative, position143)
			}
			return true
		l142:
			position, tokenIndex = position142, tokenIndex142
			return false
		},
		/* 23 decimal_point <- <'.'> */
		nil,
		/* 24 textdata <- <([a-z] / [A-Z] / [0-9] / ' ' / '!' / '#' / '$' / '&' / '%' / '*' / '+' / '-' / '.' / '/' / ':' / ';' / [<->] / '?' / '\\' / '^' / '_' / '`' / '{' / '|' / '}' / '~')> */
		func() bool {
			position145, tokenIndex145 := position, tokenIndex
			{
				position146 := position
				{
					position147, tokenIndex147 := position, tokenIndex
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l148
					}
					position++
					goto l147
				l148:
					position, tokenIndex = position147, tokenIndex147
					if c := buffer[position]; c < rune('A') || c > rune('Z') {
						goto l149
					}
					position++
					goto l147
				l149:
					position, tokenIndex = position147, tokenIndex147
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l150
					}
					position++
					goto l147
				l150:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune(' ') {
						goto l151
					}
					position++
					goto l147
				l151:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('!') {
						goto l152
					}
					position++
					goto l147
				l152:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('#') {
						goto l153
					}
					position++
					goto l147
				l153:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('$') {
						goto l154
					}
					position++
					goto l147
				l154:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('&') {
						goto l155
					}
					position++
					goto l147
				l155:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('%') {
						goto l156
					}
					position++
					goto l147
				l156:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('*') {
						goto l157
					}
					position++
					goto l147
				l157:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('+') {
						goto l158
					}
					position++
					goto l147
				l158:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('-') {
						goto l159
					}
					position++
					goto l147
				l159:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('.') {
						goto l160
					}
					position++
					goto l147
				l160:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('/') {
						goto l161
					}
					position++
					goto l147
				l161:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune(':') {
						goto l162
					}
					position++
					goto l147
				l162:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune(';') {
						goto l163
					}
					position++
					goto l147
				l163:
					position, tokenIndex = position147, tokenIndex147
					if c := buffer[position]; c < rune('<') || c > rune('>') {
						goto l164
					}
					position++
					goto l147
				l164:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('?') {
						goto l165
					}
					position++
					goto l147
				l165:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('\\') {
						goto l166
					}
					position++
					goto l147
				l166:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('^') {
						goto l167
					}
					position++
					goto l147
				l167:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('_') {
						goto l168
					}
					position++
					goto l147
				l168:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('`') {
						goto l169
					}
					position++
					goto l147
				l169:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('{') {
						goto l170
					}
					position++
					goto l147
				l170:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('|') {
						goto l171
					}
					position++
					goto l147
				l171:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('}') {
						goto l172
					}
					position++
					goto l147
				l172:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('~') {
						goto l145
					}
					position++
				}
			l147:
				add(ruletextdata, position146)
			}
			return true
		l145:
			position, tokenIndex = position145, tokenIndex145
			return false
		},
		/* 26 Action0 <- <{ p.pushArray() }> */
		nil,
		/* 27 Action1 <- <{ p.popArray() }> */
		nil,
		/* 28 Action2 <- <{ p.pushArray() }> */
		nil,
		/* 29 Action3 <- <{ p.popArray() }> */
		nil,
		nil,
		/* 31 Action4 <- <{ p.addElement(buffer[begin:end]) }> */
		nil,
		/* 32 Action5 <- <{ p.addElement(buffer[begin:end]) }> */
		nil,
		/* 33 Action6 <- <{ p.addElement(buffer[begin:end]) }> */
		nil,
	}
	p.rules = _rules
	return nil
}
