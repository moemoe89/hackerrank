# https://www.hackerrank.com/challenges/beautiful-days-at-the-movies/problem?isFullScreen=true

def beautifulDays(i, j, k):
    count = 0

    while i <= j:
        # short code
        # reverse = int(str(i)[::-1])
        reverse = reverse_func(i)

        if (i - reverse) % k == 0:
            count += 1

        i += 1

    return count

def reverse_func(n):
    reverse_n = 0

    while n > 0:
        reverse_n = reverse_n * 10 + n%10

        n //= 10

    return reverse_n

if __name__ == '__main__':
    print(beautifulDays(20, 23, 6))
