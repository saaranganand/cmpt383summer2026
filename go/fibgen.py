# fibgen.py

# // generate the Fibonacci numbers one at a time
# func fibgen() chan int {
# 	ch := make(chan int)
# 	go func() {
# 		a, b := 1, 1
# 		for { // infinite loop
# 			ch <- a // blocks here until a is removed from channel
# 			a, b = b, a+b
# 		}
# 	}() // note the () here: this is a function call
# 	return ch
# }

def fibgen():
    a, b = 1, 1
    while True:
        yield a
        a, b = b, a + b

def testFib():
    fib = fibgen()
    for i in range(10):
        print(next(fib))

testFib()

# func TestFib() {
# 	// every time <-nextFib is called, the next Fibonacci number is taken from
# 	// the channel
# 	nextFib := fibgen()
# 	for i := 0; i < 10; i++ {
# 		fmt.Println(<-nextFib)
# 	}
# }