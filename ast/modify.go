package ast

type ModifierFunc func(Node) Node

func Modify(node Node, modifier ModifierFunc) Node {
	switch node := node.(type) {
	case *Program:
		p := &Program{}
		for _, stmt := range node.Statements {
			p.Statements = append(p.Statements, Modify(stmt, modifier).(Statement))
		}
		return modifier(p)

	case *FunctionLiteral:
		fn := &FunctionLiteral{Token: node.Token}
		for _, param := range node.Parameters {
			fn.Parameters = append(fn.Parameters, Modify(param, modifier).(*Identifier))
		}
		fn.Body = Modify(node.Body, modifier).(*BlockStatement)
		return modifier(fn)
	case *ArrayLiteral:
		arr := &ArrayLiteral{Token: node.Token}
		for _, elem := range node.Elements {
			arr.Elements = append(arr.Elements, Modify(elem, modifier).(Expression))
		}

		return modifier(arr)
	case *CallExpression:
		ce := &CallExpression{Token: node.Token}
		ce.Function = Modify(node.Function, modifier).(Expression)
		for _, arg := range node.Arguments {
			ce.Arguments = append(ce.Arguments, Modify(arg, modifier).(Expression))
		}

		return modifier(ce)
	case *HashLiteral:
		hash := &HashLiteral{Token: node.Token}

		pairs := make(map[Expression]Expression)
		for key, val := range node.Pairs {
			newKey := Modify(key, modifier).(Expression)
			newVal := Modify(val, modifier).(Expression)
			pairs[newKey] = newVal
		}

		hash.Pairs = pairs
		return modifier(hash)
	case *IndexExpression:
		iexpr := &IndexExpression{Token: node.Token}
		iexpr.Left, _ = Modify(node.Left, modifier).(Expression)
		iexpr.Index, _ = Modify(node.Index, modifier).(Expression)

		return modifier(iexpr)
	case *PrefixExpression:
		pexpr := &PrefixExpression{Token: node.Token, Operator: node.Operator}
		pexpr.Right, _ = Modify(node.Right, modifier).(Expression)

		return modifier(pexpr)
	case *InfixExpression:
		iexpr := &InfixExpression{Token: node.Token, Operator: node.Operator}
		iexpr.Left, _ = Modify(node.Left, modifier).(Expression)
		iexpr.Right, _ = Modify(node.Right, modifier).(Expression)

		return modifier(iexpr)
	case *IfExpression:
		ifexpr := &IfExpression{Token: node.Token}
		ifexpr.Condition, _ = Modify(node.Condition, modifier).(Expression)
		ifexpr.Consequence, _ = Modify(node.Consequence, modifier).(*BlockStatement)
		ifexpr.Alternative, _ = Modify(node.Alternative, modifier).(*BlockStatement)

		return modifier(ifexpr)
	case *BlockStatement:
		block := &BlockStatement{Token: node.Token}
		for _, stmt := range node.Statements {
			block.Statements = append(block.Statements, Modify(stmt, modifier).(Statement))
		}

		return modifier(block)
	case *ReturnStatement:
		rstmt := &ReturnStatement{Token: node.Token}
		rstmt.ReturnValue = Modify(node.ReturnValue, modifier).(Expression)

		return modifier(rstmt)
	case *LetStatement:
		ls := &LetStatement{Token: node.Token}
		ls.Value = Modify(node.Value, modifier).(Expression)

		return modifier(ls)
	case *ExpressionStatement:
		estmt := &ExpressionStatement{Token: node.Token}
		estmt.Expression, _ = Modify(node.Expression, modifier).(Expression)

		return modifier(estmt)
	default:
		return node
	}

}
