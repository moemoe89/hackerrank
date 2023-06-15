# https://www.hackerrank.com/challenges/anagram/problem?isFullScreen=true

def anagram(s):
    n = len(s)

    if n % 2 != 0:
        return -1

    m = {}

    for c in s[:n//2]:
        m[c] = m.get(c, 0) + 1

    for c in s[n//2:]:
        m[c] = m.get(c, 0) - 1

    count = 0

    for value in m.values():
        if value > 0:
            count += value

    return count

if __name__ == '__main__':
    print(anagram('xaxbbbxx'))