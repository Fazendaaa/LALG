---
title: Relatório Trabalho Dois - Análise Sintática
author:
    - Camila Stenico
    - USP 8530952
    - Lucas De Almeida Carotta
    - USP 8598732
---

# Introdução
Este trabalho é uma continuação do **Trabalho Um - Analisador Léxico**, onde a ideia era desenvolver um analisador léxico para a linguagem teórica LALG. O trabalho atual expande isso, implementado um analisador sintátio para a linguagem utilizando o YACC.

## YACC
O tutorial apresentado em [referência](#refer%C3%AAncia) para um primeiro contato com o **YACC** utiliza o padrão **POSIX**, todavia após ver exemplos de códigos na internet de aborgagens diferentes foi decidido não seguir esse padrão devido há baixa flexibiliade dele.

# Utilização
O trabalho vem acompanhado the um arquivo `Makefile` e foi compilado em testes em um ambiente Linux. Para compilá-lo basta rodar no terminal:

```shell
$ make
```

Logo em seguida o programa será gerado, para excutá-lo:

```shell
$ ./output
```

## Dockerfile
Caso queira utilizar a imagem do Docker provida basta "compilar" o `Dockerfile` provido para simular o ambiente de desenvolvimento do trablho.

# Referência

- [Lex & Yacc](https://www.epaperpress.com/lexandyacc/)
- [Lexical Scanning in Go - Rob Pike](https://youtu.be/HxaD_trXwRE)
- [Write text parsers with yacc and lex](https://developer.ibm.com/tutorials/au-lexyacc/)
- [Writing Your Own Toy Compiler](https://gnuu.org/2009/09/18/writing-your-own-toy-compiler/)
