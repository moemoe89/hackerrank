# https://www.hackerrank.com/challenges/camelcase/problem?isFullScreen=true

def camelcase(s):
    count = 1
    
    for c in s:
        # another approach
        # if c >= 'A' and c <= 'Z':
        if c.isupper():
            count += 1
    
    return count

if __name__ == '__main__':
    print(camelcase('saveChangesInTheEditor'))
