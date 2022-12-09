data = open('input', 'r').readline()

def find_unique(marker):
    for i in range(len(data)):
        if len(set(data[i:i+marker])) == marker:
            print(i + marker)
            break

# Part 1
find_unique(4)
# Part 2
find_unique(14)