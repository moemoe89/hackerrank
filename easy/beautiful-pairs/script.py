# https://www.hackerrank.com/challenges/beautiful-pairs/problem?isFullScreen=true

def beautifulPairs(A, B):
    m = {}

    for v in A:
        m[v] = m.get(v, 0) + 1

    count = 0

    for v in B:
        if v in m and m[v] > 0:
            m[v] -= 1
            count += 1

    if count == len(A):
        return count - 1
    else:
        return count + 1

if __name__ == '__main__':
    print(beautifulPairs([10, 11, 12, 5, 14], [8, 9, 11, 11, 5]))
