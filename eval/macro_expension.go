package eval

import (
	"monkey/ast"
	"monkey/object"
)

func DefineMacros(program *ast.Program, env *object.Environment) {
	definitions := []int{}

	for i, stmt := range program.Statements {
		if isMacroDefinition(stmt) {
			addMacro(stmt, env)
			definitions = append(definitions, i)
		}
	}

	for i := len(definitions) - 1; i >= 0; i = i - 1 {
		definitionIndex := definitions[i]
		program.Statements = append(
			program.Statements[:definitionIndex],
			program.Statements[definitionIndex+1:]...,
		)
	}
}

func ExpandMacros(program *ast.Program, env *object.Environment) ast.Node {
	return ast.Modify(program, func(node ast.Node) ast.Node {
		callExpr, ok := node.(*ast.CallExpression)
		if !ok {
			return node
		}

		macro, ok := isMacroCall(callExpr, env)
		if !ok {
			return node
		}

		args := quoteArgs(callExpr)
		evalEnv := extendMacroEnv(macro, args)

		evaluated := Eval(macro.Body, evalEnv)
		quote, ok := evaluated.(*object.Quote)
		if !ok {
			panic("We only support returning AST-Nodes from macros")
		}

		return quote.Node
	})
}

func isMacroCall(expr *ast.CallExpression, env *object.Environment) (*object.Macro, bool) {
	ident, ok := expr.Function.(*ast.Identifier)
	if !ok {
		return nil, false
	}

	obj, ok := env.Get(ident.Value)
	if !ok {
		return nil, false
	}

	macro, ok := obj.(*object.Macro)
	if !ok {
		return nil, false
	}

	return macro, true
}

func quoteArgs(expr *ast.CallExpression) []*object.Quote {
	args := []*object.Quote{}

	for _, arg := range expr.Arguments {
		args = append(args, &object.Quote{Node: arg})
	}

	return args
}

func extendMacroEnv(macro *object.Macro, args []*object.Quote) *object.Environment {
	extendedEnv := object.NewEnclosedEnvironment(macro.Env)

	for idx, arg := range macro.Parameters {
		extendedEnv.Set(arg.Value, args[idx])
	}

	return extendedEnv
}

func isMacroDefinition(stmt ast.Statement) bool {
	letStatement, ok := stmt.(*ast.LetStatement)
	if !ok {
		return false
	}

	_, ok = letStatement.Value.(*ast.MacroLiteral)
	return ok
}

func addMacro(stmt ast.Statement, env *object.Environment) {
	letStatement, _ := stmt.(*ast.LetStatement)
	macroLiteral, _ := letStatement.Value.(*ast.MacroLiteral)

	macro := &object.Macro{
		Parameters: macroLiteral.Parameters,
		Body:       macroLiteral.Body,
		Env:        env,
	}

	env.Set(letStatement.Name.Value, macro)
}
