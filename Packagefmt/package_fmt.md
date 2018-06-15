####Package fmt

* ዝግጅት
    * varን በመጠቀም
        * የzero ዋጋ
    * short declaration operator
    https://play.golang.org/p/oLxC8gHE0D 
    * verbs - format specifiers (ቅርጸት መወሰኛ)
        * %v - የተባራይ ዋጋ በነባሪ ቅርጸት (value in a default format)
        * %T - የዋጋ አይነት (a Go-syntax representation of the type of the value)
    * escaped characters like \n or \t
    https://golang.org/ref/spec#Rune_literals 

####Format printing

* group #1 - general printing to standard out (ወደ መደበኛ ማውጫ ማተም)
[func Print(a ...interface{}) (n int, err error)](https://golang.org/pkg/fmt/#Print)
[func Printf(format string, a ...interface{}) (n int, err error)](https://golang.org/pkg/fmt/#Printf)
[func Println(a ...interface{}) (n int, err error)](https://golang.org/pkg/fmt/#Println)
* group #2 = printing to a string which you can assign to a variable (ወደ string ማተም)
[func Sprint(a ...interface{}) string](https://golang.org/pkg/fmt/#Sprint)
[func Sprintf(format string, a ...interface{}) string](https://golang.org/pkg/fmt/#Sprintf)
[func Sprintln(a ...interface{}) string](https://golang.org/pkg/fmt/#Sprintln)
* group #3 - printing to a file or a web server’s response (ወደ ፋይል ለማተም ወይም የድር አገልጋይ መልስ ሲሰጥ)
[func Fprint(w io.Writer, a ...interface{}) (n int, err error)](https://golang.org/pkg/fmt/#Fprint)
[func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)](https://golang.org/pkg/fmt/#Fprintf)
[func Fprintln(w io.Writer, a ...interface{}) (n int, err error)](https://golang.org/pkg/fmt/#Fprintln)
