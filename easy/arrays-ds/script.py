# https://www.hackerrank.com/challenges/arrays-ds/problem?isFullScreen=true

def reverseArray(a):
    # short code
    # return reversed(a)

    i, j = 0, len(a)-1

    while i < j:
        a[i], a[j] = a[j], a[i]
        i += 1
        j -= 1

    return a

if __name__ == '__main__':
    print(reverseArray([1, 2, 3]))
