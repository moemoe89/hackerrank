# https://www.hackerrank.com/challenges/apple-and-orange/problem?isFullScreen=true

def countApplesAndOranges(s, t, a, b, apples, oranges):
    appleFals = 0

    for apple in apples:
        if apple + a >= s and apple + a <= t:
            appleFals += 1

    orangeFalss = 0

    for orange in oranges:
        if orange + b >= s and orange + b <= t:
            orangeFalss += 1

    print(appleFals)
    print(orangeFalss)

if __name__ == '__main__':
    print(countApplesAndOranges(5, 15, 3, 2, [-2, 2, 1], [5, -6]))
