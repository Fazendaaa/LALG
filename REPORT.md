---
title: Relatório Trabalho Dois - Análise Sintática
author:
    - Camila Stenico
    - USP 8530952
    - Lucas De Almeida Carotta
    - USP 8598732
---

# Introdução

Este trabalho é uma continuação do _Trabalho Um - Analisador Léxico_, onde a ideia era desenvolver um analisador léxico para a linguagem teórica **LALG**. O trabalho atual expande isso, reimplementado a parte do analisador léxico e acrescentando o analisador sintátio para a linguagem.

## YACC

O tutorial apresentado em [referência](#refer%C3%AAncia) para um primeiro contato com o **YACC** utiliza o padrão _POSIX_, todavia após ver exemplos de códigos na internet de aborgagens diferentes foi decidido não seguir esse padrão devido há baixa flexibiliade dele. Além disso, mesmo se o padrão não fosse utilizado, como um dos membros do grupo estava implementando uma linguagem no seu tempo livre -- [TypeR](https://github.com/Fazendaaa/TypeR) --, foi decidido reaproveitar o código, apenas mudando a sintaxe para a da gramática do LALG.

# Utilização

Como o trabalho foi implementado em [Go](https://golang.org/) puro -- ou seja, sem nenhum uso de pacotes de terceiros --, para poder ser rodado é necessário que o mesmo se encontre instalado na máquina.

O trabalho vem acompanhado the um arquivo `Makefile` e foi compilado em testes em um ambiente Arch Linux. Para rodar, basta abrir o terminal na pasta raíz do projeto e:

```shell
$ make
```

Primeiramentes os testes descritos nos arquivos _*_test.go_ serão rodados para ver se a aplicação se comporta como o esperado:

```shell
ok  	_/path/LALG/src/lexer
ok  	_/path/LALG/src/ast
ok  	_/path/LALG/src/parser
```

Logo em seguida um "REPL" será iniciado para comandos serem digitados para serem analisádos:

```shell
Hello USERNAME! This is LALG programming language!
Fell free to type in commands
To exit, just type Ctrl + C
>> _
```

Caso ocorra tudo sem maiores problemas, o própiro comando será retornado. Em alguns casos, em operações aninhadas, a operação em si será encapsulada em parentesis, mostrando a ordem que a árvore de sintaxe abstrata irá percorrer para verificar depois a expressão:

```shell
>> var foo: integer := 1;
var foo: integer := 1;
>> _
```

E caso algum erro tenha acontecido, acontecerá um report inforando sobre ele:

```shell
>> var bar := 2;
	Expected next token to be :, got ':=' instead
	no prefix parse function for ':=' was found
>> _
```

# Referência

- [Lex & Yacc](https://www.epaperpress.com/lexandyacc/)
- [Lexical Scanning in Go - Rob Pike](https://youtu.be/HxaD_trXwRE)
- [Write text parsers with yacc and lex](https://developer.ibm.com/tutorials/au-lexyacc/)
- [Writing Your Own Toy Compiler](https://gnuu.org/2009/09/18/writing-your-own-toy-compiler/)
