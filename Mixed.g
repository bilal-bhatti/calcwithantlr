grammar Mixed;

start
    :   expr EOF
    ;

expr
    :   left=expr op=('*'|'/') right=expr                  # MulDiv
    |   left=expr op=('+'|'-') right=expr                  # AddSub
    |   numerator=NUM '/' denominator=NUM                  # Fraction
    |   whole=NUM '_' numerator=NUM '/' denominator=NUM    # Mixed
    |   value=NUM                                          # Number
    ;

OP_ADD: '+';
OP_SUB: '-';
OP_MUL: '*';
OP_DIV: '/';

NUM   : [0-9]+;
WS    :  [ \t\r\n] -> skip;