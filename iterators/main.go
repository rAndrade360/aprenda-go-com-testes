package iterators

const repeatQtd = 4

func Repeat(iterator string) string {
	var repeat string
	for i := 0; i < repeatQtd; i++ {
		repeat += iterator
	}

	return repeat
}

func main() {

}
