# https://www.hackerrank.com/challenges/between-two-sets/problem?isFullScreen=true

def getTotalX(a, b):
    output = 0
    
    lastA = a[len(a)-1]
    firstB = b[0] + 1
    
    for i in range(lastA, firstB):
        is_factor_a = True
        
        for v in a:
            if i % v != 0:
                is_factor_a = False
                break
        
        if is_factor_a == False:
            continue
        
        is_factor_b = True
        
        for v in b:
            if v % i != 0:
                is_factor_b = False
                break
            
        if is_factor_b == True:
            output += 1
    
    return output

if __name__ == '__main__':
    print(getTotalX([2, 4], [16, 32, 96]))
