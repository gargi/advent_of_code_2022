file = open("input").read().strip()

occupied = set()
max_y = 0
lines = file.split("\n")
for line in lines:
    points = line.split("->")
    for i in range(1, len(points)):
        x, y = points[i].split(",")
        x, y = int(x), int(y)
        prev_x, prev_y = points[i-1].split(",")
        prev_x, prev_y = int(prev_x), int(prev_y)
        max_y = max(max_y, int(prev_y))
        if x == prev_x:
            for y_coord in range(min(y, prev_y), max(y, prev_y)+1):
                occupied.add((x, y_coord))
        else:
            for x_coord in range(min(x, prev_x), max(x, prev_x)+1):
                occupied.add((x_coord, y))
max_y += 1

units = 0
while True:
    sand_x = 500
    sand_y = 0
    while True:
        if sand_y >= max_y: # abyss
           print("Part 1: ", units)
           exit(0)
        if (sand_x, sand_y+1) not in occupied: # continues straight down
            sand_y += 1
        elif (sand_x-1, sand_y+1) not in occupied: # left
            sand_x -= 1
            sand_y += 1
        elif (sand_x+1, sand_y+1) not in occupied: # right
            sand_x += 1
            sand_y += 1
        else:
            occupied.add((sand_x, sand_y))  # rest
            units += 1
            break



