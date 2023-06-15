# https://www.hackerrank.com/challenges/alternating-characters/problem?isFullScreen=true

def alternatingCharacters(s):
    count = 0

    for i in range(len(s)):
        if i == 0:
            continue

        if s[i] == s[i-1]:
            count += 1

    return count

if __name__ == '__main__':
    print(alternatingCharacters('AAABBB'))
