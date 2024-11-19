package week4

/*
//https://www.hackerrank.com/challenges/three-month-preparation-kit-kangaroo/problem

x - point where kangaroos meet
x1+xv1 = x2+xv2
x1-x2 = xv2 - xv1
x1-x2 = x(v2-v1)
(x1-x2)/(v2-v1) = x => v2-v1!= 0, x>0
*/
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v2-v1 == 0 {
		return "NO"
	}
	if v2-v1 < 0 && x1-x2 > 0 {
		return "NO"
	}
	if x1-x2 < 0 && v2-v1 > 0 {
		return "NO"
	}
	if (x1-x2)%(v2-v1) == 0 {
		return "YES"
	}
	return "NO"
}
