# https://www.hackerrank.com/challenges/beautiful-triplets/problem?isFullScreen=true

def beautifulTriplets(d, arr):
    m = {}

    for v in arr:
        m[v] = m.get(v, True)

    count = 0

    for v in arr:

        if (v+d in m) and (v+d*2 in m):
            count += 1

    return count

if __name__ == '__main__':
    print(beautifulTriplets([1, 2, 4, 5,7, 8, 10]))
