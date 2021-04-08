<!DOCTYPE html>
<h1>📚Golang formatting📚</h1>
<h1>Printing The verbs:</h1>

<h3>%v	the value in a default format</h3>
<h3>when printing structs, the plus flag (%+v) adds field names</h3>
<h3>%#v	a Go-syntax representation of the value</h3>
<h3>%T	a Go-syntax representation of the type of the value</h3>
<h3>%%	a literal percent sign; consumes no value</h3>

<h1>Boolean:</h1>

<h3>%t	the word true or false</h3>

<h1> INT: </h1>
<h3>%b	base 2</h3>
<h3>%c	the character represented by the corresponding Unicode code point</h3>
<h3>%d	base 10</h3>
<h3>%o	base 8</h3>
<h3>%O	base 8 with 0o prefix</h3>
<h3>%q	a single-quoted character literal safely escaped with Go syntax.</h3>
<h3>%x	base 16, with lower-case letters for a-f</h3>
<h3>%X	base 16, with upper-case letters for A-F</h3>
<h3>%U	Unicode format: U+1234; same as "U+%04X"</h3>

<h1>Float:</h1>

<h3>%b	decimalless scientific notation with exponent a power of two,in the manner of strconv.FormatFloat with the 'b' format,e.g. -123456p-78</h3>
<h3>%e	scientific notation, e.g. -1.234456e+78</h3>
<h3>%E	scientific notation, e.g. -1.234456E+78</h3>
<h3>%f	decimal point but no exponent, e.g. 123.456</h3>
<h3>%F	synonym for %f</h3>
<h3>%g	%e for large exponents, %f otherwise. Precision is discussed below.</h3>
<h3>%G	%E for large exponents, %F otherwise</h3>
<h3>%x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20</h3>
<h3>%X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20</h3>

<h1>String and slice of bytes (treated equivalently with these verbs):</h1>

<h3>%s	the uninterpreted bytes of the string or slice</h3>
<h3>%q	a double-quoted string safely escaped with Go syntax</h3>
<h3>%x	base 16, lower-case, two characters per byte</h3>
<h3>%X	base 16, upper-case, two characters per byte</h3>
<h3>_____________________________________________________________________________________________</h3>
<h3>bool:                    %t</h3>
<h3>int, int8 etc.:          %d</h3>
<h3>uint, uint8 etc.:        %d, %#x if printed with %#v</h3>
<h3>float32, complex64, etc: %g</h3>
<h3>string:                  %s</h3>
<h3>chan:                    %p</h3>
<h3>pointer:                 %p</h3>

<h1> Sintaxe</h1>
<h3>struct:             {field0 field1 ...}</h3>
<h3>array, slice:       [elem0 elem1 ...]</h3>
<h3>maps:               map[key1:value1 key2:value2 ...]</h3>
<h3>pointer to above:   &{}, &[], &map[]</h3>

<h1> Precision</h1>
<h3>%f     default width, default precision</h3>
<h3>%9f    width 9, default precision</h3>
<h3>%.2f   default width, precision 2</h3>
<h3>%9.2f  width 9, precision 2</h3>
<h3>%9.f   width 9, precision 0</h3>

<h3>Link to access the official documentation, all this was taken from there:https://golang.org/pkg/fmt/</h3>
</html>
