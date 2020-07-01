// := --> declaration and assignment, = --> assignment only
// Only iterator available is for
package iteration

func Repeat(character string) string{
	var str string
	for i:=0; i<5; i++{
		str += character
	}
	return str
}
