# https://www.hackerrank.com/challenges/append-and-delete/problem?isFullScreen=true

def appendAndDelete(s, t, k):
    lenS = len(s)
    lenT = len(t)

    if lenS+lenT <= k:
        return 'Yes'

    samePrefix = 0
    while samePrefix < lenS and samePrefix < lenT and s[samePrefix] == t[samePrefix]:
        samePrefix += 1

    remaingS = lenS - samePrefix
    remaingT = lenT - samePrefix

    remaining = k - (remaingS+remaingT)
    if remaining >= 0 and remaining % 2 == 0:
        return 'Yes'

    return 'No'

if __name__ == '__main__':
    print(appendAndDelete('hackerhappy', 'hackerrank', 9))
