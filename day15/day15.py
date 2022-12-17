from parse import parse

y=2000000

file = open("input").read()
no_signal = set()
lines = file.split("\n")
for line in lines:
    sensor_x, sensor_y, beacon_x, beacon_y = parse("Sensor at x={:d}, y={:d}: closest beacon is at x={:d}, y={:d}", line)
    mh_distance = abs(sensor_x - beacon_x) + abs(sensor_y - beacon_y)
    for x in range (sensor_x - mh_distance, sensor_x + mh_distance + 1):
            if abs(sensor_x - x) + abs(sensor_y - y) <= mh_distance:
                no_signal.add((x,y))
    no_signal.discard((beacon_x, beacon_y))
print(len(no_signal))

