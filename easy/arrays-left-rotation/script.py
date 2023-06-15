# https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem?isFullScreen=true

def rotLeft(a, d):
    d = d % len(a)
    return a[d:] + a[:d]

if __name__ == '__main__':
    print(rotLeft([1, 2, 3], 10))
