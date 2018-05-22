#### የGo   ቋንቋ  መሰረታዊ መዋቅር:
- package main - የፓኬጁ ስም 
- import fmt ፡ ሌላ ፓኬጅ ኢምፖርት በማድረግ መጠቀም ፣ የአናጺ ምሳሌ
- func main ፡ የprogram መግቢያና ፡ መውጫ ነው ። 
- ኮምፒውተራችን የጻፍነውን ኮድ መተግበር የሚጀምረው ከዚህ ፈንክሽን ጀምሮ ነው ፣ ልክ የሄ ፈንክሽን ሲያልቅ ፕሮግራማችኝ ስራውን ይጨርሳል
- የGo መሰረታዊ ወይንም ቁልፍ ቃላት ፡ golang spec key words ፡ package func import
- fmt ፡ golang.org -> packages ወይንም godoc.org / fmt
- የመጀመሪያ ፊደላቸው ካፒታል የሆነ ፈንክሽኖች exported/visible የሚሆኑ ናቸው ፣ ከሌላ ቦታ መጠራት የሚችሉ ማለት ነው (Println)
variadic parameters ምን ማለት ነው ፣ ያልተገደበ(unlimited) የግቤት(parameter) ብዛት  ይወክላል
the “...<some type>” is how we signify a variadic parameter
the type “interface{}” is the empty interface
ምንኛውም/ሁሉም ነገር  is of type “interface{}”
so the parameter “...interface{}” means “give me as many arguments of any type as you’d like
የማይፈለጉ የፈንክሽን  returns ማጥፋት ፡ use the “_” underscore character to throw away returns
Go ጥቅም ላይ ያልዋለ ቫሪያብልን አይፈቅድም 
you can’t have unused variables in your code
this is code pollution
the compiler doesn’t allow it
ፈንክሽን የመጥሪያ ቅርጸት  notation  Go ፡ package.Identifier
ex: fmt.Println()     የፓኬጅስም.መለያስም
we would read that: “from package fmt I am using the Println func”
መለያ ስም ፡ የ ቫሪያብል ፣ የፈንክሽን ወይም የ ኮንስታንት ስም ሊሆን ይችላል
an identifier is the name of the variable, constant, func
Packages፡ አስቀድመው የተዘጋጀ ኮድ … ኢምፖርት ይደረጋሉ ፣ ሌላ ኮድ ለመጻፍ የምንዋሰው ማለት ነው ፡ 
ምሳሌ ፡ አናጺ መዶሻ አያመርትም
code that is already written which you can use -›Imports
ዋና መልክት ፡ package main እና func main
