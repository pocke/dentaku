%{
package parser

import (
  "github.com/pocke/dentaku/ast"
)

type Token struct {
	Token   int
	Literal string
}

%}


%union{
  token Token
  expr ast.Expression
}

%type<expr> program
%type<expr> expr add_expr sub_expr multi_expr div_expr paren_expr
%token<token> LITERAL

%left '+' '-'
%left '*' '/'

%%

program
  : expr
  {
    $$ = $1
    yylex.(*Lexer).result = $$
  }

expr
	: LITERAL
	{
		$$ = ast.Literal{Literal: $1.Literal}
	}
	| add_expr
	| sub_expr
	| multi_expr
	| div_expr
  | paren_expr

add_expr
	: expr '+' expr
	{
		$$ = ast.BinOpExpr{Left: $1, Operator: '+', Right: $3}
	}

sub_expr
	: expr '-' expr
	{
		$$ = ast.BinOpExpr{Left: $1, Operator: '-', Right: $3}
	}

multi_expr
	: expr '*' expr
	{
		$$ = ast.BinOpExpr{Left: $1, Operator: '*', Right: $3}
	}

div_expr
	: expr '/' expr
	{
		$$ = ast.BinOpExpr{Left: $1, Operator: '/', Right: $3}
	}

paren_expr
	: '(' expr ')'
	{
		$$ = $2
	}

%%
