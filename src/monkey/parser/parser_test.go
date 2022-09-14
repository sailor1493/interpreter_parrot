package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10; let foobar = 2323232;
	`

	testProgram := makeProgram(t, input)

	if testProgram == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	if len(testProgram.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(testProgram.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := testProgram.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return add(15);
	`

	testProgram := makeProgram(t, input)

	if len(testProgram.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(testProgram.Statements))
	}

	for _, stmt := range testProgram.Statements {
		if !testReturnStatement(t, stmt) {
			return
		}
	}
}

func TestIdentifierExpressions(t *testing.T) {
	input := "foobar; barfoo"

	testProgram := makeProgram(t, input)

	if len(testProgram.Statements) != 2 {
		t.Fatalf("testProgram has not enough statements. got=%d", len(testProgram.Statements))
	}
	tests := []struct {
		expectedValue   string
		expectedLiteral string
	}{
		{"foobar", "foobar"},
		{"barfoo", "barfoo"},
	}
	for i, stmt := range testProgram.Statements {
		if !testIdentifierStatement(t, stmt, tests[i].expectedValue, tests[i].expectedLiteral) {
			return
		}
	}
}

// Statement Checking Internal Functions

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let', got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}

func testReturnStatement(t *testing.T, stmt ast.Statement) bool {
	returnStmt, ok := stmt.(*ast.ReturnStatement)
	if !ok {
		t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
		return false
	}
	if stmt.TokenLiteral() != "return" {
		t.Errorf("stmt.TokenLiteral not 'return', got=%q", returnStmt.TokenLiteral())
		return false
	}
	return true
}

func testIdentifierStatement(t *testing.T, stmt ast.Statement, expectedValue string, expectedLiteral string) bool {
	expStmt, ok := stmt.(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt is not *ast.ExpressionStatement. got=%T", stmt)
		return false
	}
	return testIdentifierExpression(t, &expStmt.Expression, expectedValue, expectedLiteral)
}

// Expressions Checking Internal Functions

func testIdentifierExpression(t *testing.T, expr_ptr *ast.Expression, expectedValue string, expectedLiteral string) bool {
	expression := *expr_ptr
	ident, ok := expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression is not *ast.Identifier. got=%T", expression)
		return false
	}
	if ident.Value != expectedValue {
		t.Errorf("ident.Value not %s. got=%s", expectedValue, ident.Value)
		return false
	}
	if ident.TokenLiteral() != expectedLiteral {
		t.Errorf("ident.TokenLiteral not %s. got=%s", expectedLiteral, ident.TokenLiteral())
		return false
	}
	return true
}

// Error Checking Internal Functions

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func makeProgram(t *testing.T, input string) *ast.Program {

	testLexer := lexer.New(input)
	testParser := New(testLexer)
	checkParserErrors(t, testParser)
	return testParser.ParseProgram()
}
