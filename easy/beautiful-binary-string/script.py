# https://www.hackerrank.com/challenges/beautiful-binary-string/problem?isFullScreen=true

def beautifulBinaryString(b):
    count, i = 0, 0

    while i < len(b)-2:
        if b[i:i+3] == '010':
            count += 1
            i += 3
        else:
            i += 1

    return count

if __name__ == '__main__':
    print(beautifulBinaryString('0101010'))
