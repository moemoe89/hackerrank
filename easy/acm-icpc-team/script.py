# https://www.hackerrank.com/challenges/acm-icpc-team/problem?isFullScreen=true

def acmTeam(topic):
    maxTopics = 0
    m = {}

    for i in range(len(topic)-1):
        for j in range (i + 1, len(topic)):
# iterates approach
#             topics = 0
#             for k in range(len(topic[i])):
#                 if topic[i][k] == '1' or topic[j][k] == '1':
#                     topics += 1

            topics = bin(int(topic[i], 2) | int(topic[j], 2)).count('1')
            if topics in m:
                m[topics] += 1
            else:
                m[topics] = 1

            maxTopics = max(maxTopics, topics)

    return [maxTopics, m[maxTopics]]

if __name__ == '__main__':
    print(acmTeam(['10101', '11110', '00010']))
