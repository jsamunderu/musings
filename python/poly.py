# Horner's method
def polynomial(coef, x):
    y = 0
    for i in range(len(coef)-1, -1, -1):
        y = y * x + coef[i]
    return y

def polynomial_recursive(coef, x, n):
    if n == (len(coef)-1):
        return coef[n]
    return polynomial_recursive(coef, x, n+1) * x + coef[n]

if __name__ == '__main__':
    # evaluate 2x^2 + 5x + 7 where x = 13
    coef = [7, 5, 2]
    x = 13
    print(polynomial(coef, x))
    print(polynomial_recursive(coef, x, 0))
