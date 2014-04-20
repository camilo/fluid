%{
package parse
import ("errors"
)

var ParseTree ListNode
%}

%union{
  typ int
  pos int
  val string
  node *parseNode
}

%token <val> itemError
%token <val> itemIdentifier
%token <val> itemString
%token <val> itemNumberLit
%token <val> itemDotDot
%token <val> itemComparison
%token <val> itemLeftDelimLogic   // {%
%token <val> itemRightDelimLogic  // %}
%token <val> itemLeftDelimOutput  // {{
%token <val> itemRightDelimOutput  // }}
%token <val> itemPipe
%token <val> itemDot
%token <val> itemColon
%token <val> itemComma
%token <val> itemOpenSquare
%token <val> itemCloseSquare
%token <val> itemOpenRound
%token <val> itemCloseRound
%token <val> itemText
%token <val> itemEOF 1 // Fix this, yacc really wants the end to be 0 and when this token is given a 0 value it starts complaining
%token <val> itemSpace
%token <val> itemField
%token <val> itemNumber
%token <val> itemComplex

%type<node> template
%type<node> liquid
%type<node> outputSource
%type<node> outputData
%type<node> outputFilter

%%

template: /* empty */
        {
          ParseTree.appendNode(&parseNode{typ: -1})
        }
        | liquid
        {
          ParseTree.appendNode(&parseNode{typ: -1, left: $1})
        }
        | itemText
        {
          text := &parseNode{typ: textNode, value: $1}
          ParseTree.appendNode(&parseNode{typ: -1, left: text})
        }
        | itemText liquid
        {
          text := &parseNode{typ: textNode, value: $1}
          ParseTree.appendNode(&parseNode{-1, text, $2, nil})
        }
        | liquid itemText
        {
          text := &parseNode{typ: textNode, value: $2}
          ParseTree.appendNode(&parseNode{-1, $1, text, nil})
        }
        | template liquid
        {
          ParseTree.appendNode($2)
        }
        | template itemText
        {
          text := &parseNode{typ: textNode, value: $2}
          ParseTree.appendNode(text)
        }
        ;

 liquid: itemLeftDelimOutput outputData itemRightDelimOutput
       {
       // outputData rule returns liquidNode
       $$ = $2
       }
      ;

outputSource: itemString
            {$$ = &parseNode{typ: litNode, value: $1}}
            | itemNumberLit
            {$$ = &parseNode{typ: litNode, value: $1}}
            | itemNumber
            {$$ = &parseNode{typ: litNode, value:$1}}
            | itemIdentifier
            {$$ = &parseNode{typ: identifierNode, value: identifier{$1, ""}}}
            | itemIdentifier itemDot itemField
            {$$ = &parseNode{typ: identifierNode, value: identifier{$1, $3}}}
            ;

outputFilter: itemIdentifier
            { //fmt.Printf("OMG: %s\n", $1)
            $$=&parseNode{typ: filterListNode, value: $1}}
            | itemIdentifier itemColon outputSource
            { //fmt.Printf("OMG: %s : %s \n", $1, $3.value)
            $$=&parseNode{typ: filterListNode, value: $1}}
            | outputFilter itemPipe outputFilter
            {$$=&parseNode{typ: filterListNode, value: nil}}
            ;

outputData:  outputSource
      {
        source := &parseNode{typ: outputSourceNode, value: $1}
        $$ = &parseNode{typ: liquidNode, left: source}
      }
      | outputSource itemPipe outputFilter
      {
        source := &parseNode{typ: outputSourceNode, value: $1}
        filters := &parseNode{typ: filterListNode, value: nil}
        $$ = &parseNode{typ: liquidNode, left: source, right: filters}
      }
      ;
%%

func Parse(liquid string) (*ListNode, error) {
  if 0 == yyParse(lex("lexer", liquid)) {
    return &ParseTree, nil
  } else {
    return nil, errors.New("")
  }

}




