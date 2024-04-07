grammar Stack;

// Tokens
INT: 'int';
FLOAT: 'float';
BOOL: 'bool';
STRING: 'string';
TRUE: 'true';
FALSE: 'false';
WRITE: 'write';
READ: 'read';
IF: 'if';
ELSE: 'else';
WHILE: 'while';
ID: [a-zA-Z][a-zA-Z0-9]*;
INT_LITERAL: [0-9]+;
FLOAT_LITERAL: [0-9]+ '.' [0-9]+;
STRING_LITERAL: '"' (~["])* '"';
ASSIGN: '=';
OR: '||';
AND: '&&';
EQ: '==';
NE: '!=';
LT: '<';
GT: '>';
ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/';
MOD: '%';
NOT: '!';
SEM: ';';
LPAR: '(';
RPAR: ')';
LBRACE: '{';
RBRACE: '}';
COMMA: ',';
DOT: '.';
COMMENT: '//' ~[\r\n]* -> skip;
WHITE_SPACES: [ \t\r\n]+ -> skip;

// Rules
program: statement*;

type: INT
    | FLOAT
    | BOOL
    | STRING
    ;

literal: INT_LITERAL # intLiteral
       | FLOAT_LITERAL # floatLiteral
       | TRUE # trueLiteral
       | FALSE # falseLiteral
       | STRING_LITERAL # stringLiteral
       ;

variableList: ID (COMMA ID)*;

expressionList: expression (COMMA expression)*;

statement: SEM # emptySemStatement
         | type variableList SEM # variableStatement
         | expression SEM # expressionStatement
         | READ variableList SEM # readStatement
         | WRITE expressionList SEM # writeStatement
         | LBRACE statement* RBRACE # blockStatement
         | IF LPAR expression RPAR statement (ELSE statement)? # ifStatement
         | WHILE LPAR expression RPAR statement # whileStatement
         ;

expression: LPAR expression RPAR # parenExpression
          | op=(NOT | SUB) expression # unaryExpression
          | expression op=(LT | GT | EQ | NE) expression # comparisonExpression
          | expression op=(AND | OR) expression # logicalExpression
          | expression op=(MUL | DIV | MOD) expression # mulDivModExpression
          | expression op=(ADD | SUB) expression # addSubExpression
          | STRING_LITERAL DOT STRING_LITERAL # stringConcatExpression
          | literal # literalExpression
          | ID ASSIGN expression # assignExpression
          | ID # idExpression
          ;

