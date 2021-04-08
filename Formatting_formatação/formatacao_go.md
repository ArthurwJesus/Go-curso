
<h1> 📚Formatação de Golang📚 </h1>
<h1> Imprimindo os verbos: </h1>

<h3>% v o valor em um formato padrão </h3>
<h3> ao imprimir estruturas, o sinalizador de mais (% + v) adiciona nomes de campo </h3>
<h3>% # v uma representação da sintaxe Go do valor </h3>
<h3>% T a representação da sintaxe Go do tipo do valor </h3>
<h3> %% um sinal de porcentagem literal; não consome nenhum valor </h3>

<h1> Booleano: </h1>

<h3>% t a palavra verdadeiro ou falso </h3>

<h1> INT: </h1>
<h3>% b base 2 </h3>
<h3>% c o caractere representado pelo ponto de código Unicode correspondente </h3>
<h3>% d base 10 </h3>
<h3>% o base 8 </h3>
<h3>% O base 8 com prefixo 0o </h3>
<h3>% q um literal de caractere entre aspas simples com escape seguro com sintaxe Go. </h3>
<h3>% x base 16, com letras minúsculas para a-f </h3>
<h3>% X base 16, com letras maiúsculas para A-F </h3>
<h3>% U formato Unicode: U + 1234; igual a "U +% 04X" </h3>
<h1> Float: </h1>

<h3>% b notação científica sem casas decimais com expoente uma potência de dois, na forma de strconv.FormatFloat com o formato 'b', por exemplo. -123456p-78 </h3>
<h3>% e notação científica, por exemplo -1,234456e + 78 </h3>
<h3> notação científica% E, por exemplo -1,234456E + 78 </h3>
<h3>% f ponto decimal, mas nenhum expoente, por exemplo 123,456 </h3>
<h3>% F sinônimo para% f </h3>
<h3>% g% e para expoentes grandes,% f caso contrário. A precisão é discutida abaixo. </h3>
<h3>% G% E para expoentes grandes,% F caso contrário </h3>
<h3>% x notação hexadecimal (com potência decimal de dois expoentes), por ex. -0x1,23abcp + 20 </h3>
<h3>% X notação hexadecimal em maiúsculas, por ex. -0X1,23ABCP + 20 </h3>

<h1> String e fatia de bytes (tratados de forma equivalente a estes verbos): </h1>

<h3>% s os bytes não interpretados da string ou fatia </h3>
<h3>% q uma string entre aspas duplas com escape seguro com sintaxe Go </h3>
<h3>% x base 16, minúsculas, dois caracteres por byte </h3>
<h3>% X base 16, maiúsculas, dois caracteres por byte </h3>
<h3> _____________________________________________________________________________________________ </h3>
<h3> bool:% t </h3>
<h3> int, int8 etc .:% d </h3>
<h3> uint, uint8 etc .:% d,% # x se impresso com% # v </h3>
<h3> float32, complex64, etc:% g </h3>
<h3> string:% s </h3>
<h3> chan:% p </h3>
<h3> ponteiro:% p </h3>

<h1> Sintaxe </h1>
<h3>struct:             {field0 field1 ...}</h3>
<h3>array, slice:       [elem0 elem1 ...]</h3>
<h3>maps:               map[key1:value1 key2:value2 ...]</h3>
<h3>pointer to above:   &{}, &[], &map[]</h3>

<h1> Precisão </h1>
<h3>//a precisão é quantida de zeros depois do ponto//</h3>
<h3>% f largura padrão, precisão padrão </h3>
<h3>% 9f largura 9, precisão padrão </h3>
<h3>% .1f largura padrão, precisão 2 </h3> 
<h3>% .2f largura padrão, precisão 2 </h3>
<h3>% 9.2f largura 9, precisão 2 </h3>
<h3>% 9.f largura 9, precisão 0 </h3>

<h3> Link para acessar a documentação oficial, tudo retirado de lá, apenas traduzido: https: //golang.org/pkg/fmt/ </h3>

