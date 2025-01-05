grammar Finscript;

// Program is 1..* sendStmt
program
  : (sendStmt)+ EOF
  ;

sendStmt
  : SEND LBRACK ASSET AMOUNT RBRACK LPAREN sourceClause destinationClause RPAREN
  ;

// e.g. source = ...
sourceClause
  : SOURCE EQ sourceBlock
  ;

// e.g. destination = ...
destinationClause
  : DESTINATION EQ destBlock
  ;

sourceBlock
  // SingleAccount => e.g. @foo
  : ACCOUNT
  // or block => { sourceBlock (',' sourceBlock)* }
  | LBRACE inorderSrcList? RBRACE
  ;

inorderSrcList
  : sourceBlock (COMMA sourceBlock)*
  ;

destBlock
  : ACCOUNT
  | LBRACE allotmentDestList? RBRACE
  ;

allotmentDestList
  : allotmentClause (COMMA allotmentClause)*
  ;

allotmentClause
  // e.g. 85% to @account
  : portion 'to' ACCOUNT
  // or "remaining to @account"
  | REMAINING 'to' ACCOUNT
  ;

// portion =>  e.g. 85% or 1/3
portion
  : PERCENTAGE_PORTION_LITERAL
  | RATIO_PORTION_LITERAL
  ;

// Lexer rules
SEND         : 'send';
SOURCE       : 'source';
DESTINATION  : 'destination';
REMAINING    : 'remaining';

EQ           : '=';
LPAREN       : '(';
RPAREN       : ')';
LBRACK       : '[';
RBRACK       : ']';
LBRACE       : '{';
RBRACE       : '}';
COMMA        : ',';

ASSET        : [A-Z0-9/]+;
AMOUNT
 : [0-9]+
 | '*'
 ;

ACCOUNT
 : '@' [a-zA-Z0-9_:-]+
 ;

PERCENTAGE_PORTION_LITERAL
 : [0-9]+ ( '.' [0-9]+ )? '%'
 ;

RATIO_PORTION_LITERAL
 : [0-9]+ '/' [0-9]+
 ;

WS : [ \t\r\n]+ -> skip;

LINE_COMMENT : '//' ~[\r\n]* -> skip;
