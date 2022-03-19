# Horner's method
def polynomial(coef, x):
    y = 0
    for i in range(len(coef)-1, -1, -1):
        y = y * x + coef[i]
    return y

if __name__ == '__main__':
    # evaluate 2x^2 + 5x + 7 where x = 13
    coef = [7, 5, 2]
    x = 13
    print(polynomial(coef, x))
