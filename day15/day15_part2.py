from parse import parse

lower_limit = 0
upper_limit = 4000000

def parse_input():
    sensor_distance = []
    file = open("input").read().strip()
    no_signal = set()
    lines = file.split("\n")
    for line in lines:
        sensor_x, sensor_y, beacon_x, beacon_y = parse("Sensor at x={:d}, y={:d}: closest beacon is at x={:d}, y={:d}", line)
        mh_distance = abs(sensor_x - beacon_x) + abs(sensor_y - beacon_y)
        sensor_distance.append((sensor_x, sensor_y, mh_distance))
    return sensor_distance

def find_beacon(sensor_distance):
    for y_position in range(upper_limit+1):
        no_beacon = []
        for sd in sensor_distance:
            sensor_x, sensor_y, mh_distance = sd
            y_distance = abs(sensor_y - y_position)
            if y_distance <= mh_distance:
                remaining_val = abs(mh_distance - y_distance)
                start = max(lower_limit, sensor_x - remaining_val)
                end = min(upper_limit, sensor_x + remaining_val)
                no_beacon.append((start,end))

        x_position = lower_limit
        for start_bound, end_bound in sorted(no_beacon):
            if x_position >= start_bound:
                x_position = max(x_position, end_bound)
            else:
                return (x_position+1)*4000000 + y_position

sensor_distance = parse_input()
tuner_frequency = find_beacon(sensor_distance)
print(tuner_frequency)

