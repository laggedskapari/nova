package token

type TokenType string 

type Token struct {
  Type TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  IDENTIFIER = "IDENTIFIER"
  INT = "INT"

  ASSIGN = "="
  PLUS = "+"
  MINUS = "-"
  MODULO = "%"
  DIVIDE = "/"
  MULTIPLY = "*"
  LCHEVRON = "<"
  RCHEVRON = ">"
  BANG = "!"
  EQ = "=="
  NOT_EQ = "!="

  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"
  SQLBRACE = "["
  SQRBRACE = "]"

  FUNCTION = "FUNCTION"
  LET = "LET"
  IF = "IF"
  ELSE = "ELSE"
  RETURN = "RETURN"
  TRUE = "TRUE"
  FALSE = "FALSE"

)

var keywords = map[string]TokenType {
  "fn": FUNCTION,
  "let": LET,
  "true": TRUE,
  "false": FALSE, 
  "if": IF,
  "else": ELSE,
  "return": RETURN,
}

func LookupIdent(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENTIFIER
}
