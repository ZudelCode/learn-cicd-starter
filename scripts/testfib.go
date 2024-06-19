package fib_test

type fibTest struct {
        n        int
        expected int
}

var fibTests = []fibTest {
        {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13},
}
