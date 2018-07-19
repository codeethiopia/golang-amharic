#### String (ሐረግ?)

* [ስትሪንግ](https://golang.org/ref/spec#String_types) የ”ስትሪንግ ዋጋ ስብስቦችን” የሚወክል አይነት ነው። 

* ዋጋው የ ባይቶች (bytes) ስብስብ ነው ፣ ባዶ ሊሆንም ይችላል።

* ስትሪንግ አንዴ ከተፈጠረ በኋላ አይቀየርም (Strings are immutable: once created, it is impossible to change the contents of a string.)

* ስትሪንግ ለመፍጠር ፣ "" (double quotes) ወይም `` (backticks)

**ምሳሌ**

* በ `` የሚፈጠር ስትሪንግ [string literal](https://golang.org/ref/spec#String_literals)  ይባላል። returns እና spaces  መኖር ይችላሉ 
    
    * “\” (backslashes) ልዩ ትርጉም የለውም 

**ምሳሌ**

* ዋጋው የ ባይቶች (bytes) ስብስብ ነው ፣  sequence of bytes ወይም slice of bytes

* ባይት የ int8 ሌላ መጠሪያ (alias) ነው

* ስትሪንግን ወደ ባይት ዝርዝር (slice of bytes፣ []byte(s)) መቀየር እንችላለን።

**ምሳሌ**

    * የ [ASCII](https://en.wikipedia.org/wiki/ASCII) ኮዲንግ ስኪም ላይ እያንዳንዱ ባይት የሚወክለውን ፊደል እንይ ።

* ከዴሲማል በተጨማሪ በሄክሳዴሲማልና በ UTF-8  የስትሪንጉን የኢንኮዲንግ ዋጋ ማየት ከፈለግን የ [fmt ፓኬጅ](https://golang.org/pkg/fmt/) ውስጥ ያሉ ቅርጸት ገላጭ (formatting verbs) በድጋሚ እንይ …

```
%s	the uninterpreted bytes of the string or slice
%q	a double-quoted string safely escaped with Go syntax
%x	base 16, lower-case, two characters per byte
%X	base 16, upper-case, two characters per byte
%U	Unicode format: U+1234; same as "U+%04X"
```

Other flags:

```
+	always print a sign for numeric values;
	guarantee ASCII-only output for %q (%+q)
-	pad with spaces on the right rather than the left (left-justify the field)
#	alternate format: add leading 0 for octal (%#o), 0x for hex (%#x);
	0X for hex (%#X); suppress 0x for %p (%#p);
	for %q, print a raw (backquoted) string if strconv.CanBackquote
	returns true;
	always print a decimal point for %e, %E, %f, %F, %g and %G;
	do not remove trailing zeros for %g and %G;
	write e.g. U+0078 'x' if the character is printable for %U (%#U).
' '	(space) leave a space for elided sign in numbers (% d);
	put spaces between bytes printing strings or slices in hex (% x, % X)
0	pad with leading zeros rather than spaces;
	for numbers, this moves the padding after the sign
```

**ምሳ ** 

* Each code point is known as a rune, which is an alias for int32. Each rune is a code point in UTF-8.


#### ዋና መልእክት ፤ 

* ስትሪንግ በ double quotes ወይም በ backticks ነው የሚፈጠረው

* የGo ኮድ በ UTF-8 ነው ኢንኮድ የሚደረገው (ነገር ግን ሁሉም የስትሪንግ የUTF-8 code point ላይሆን ይችላል)

* ስትሪንግ የ ባይት ዝርዝር ነው (slice of bytes)

* Immutable፣ የስትሪንግ ዋጋ መቀየር አይችልም ፣ አዲስ ዋጋ መስጠት ብቻ ነው የምንችለው።


