# https://www.hackerrank.com/challenges/angry-professor/problem?isFullScreen=true

def angryProfessor(k, a):
    # attend = 0
    # for v in a:
    #     if v <= 0:
    #         attend += 1

    # short code
    attend = sum(1 for x in a if x <= 0)

    if attend >= k:
        return 'NO'
    else:
        return 'YES'

if __name__ == '__main__':
    print(angryProfessor(3, [-1, -3, 4, 2]))
