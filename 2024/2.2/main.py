def isSorted(level, compF):
    for i in range(1, len(level)):
        if not compF(level[i], level[i-1]):
            return False

    return True


def isSorted2(level, compF, recCount = 0):
    if recCount > 1:
        return False

    for i in range(1, len(level)):
        if not compF(level[i], level[i-1]):
            return isSorted2(level[:i] + level[i+1:], compF, recCount+1) or isSorted2(level[:i-1] + level[i:], compF, recCount+1)

    return True


file = open("input.txt", "r")
lines = file.readlines()

levels = []
for l in lines:
    levels.append([int(x) for x in l.split(" ")])

safeReportCount = 0
for l in levels:
    ascendingSorted = isSorted2(l, lambda x, y: (x < y) and (x + 3 >= y))
    descendingSorted = isSorted2(l, lambda x, y: (x > y) and (x <= y + 3))
    if ascendingSorted or descendingSorted:
        safeReportCount += 1
    else:
        print(l)

print(f"Safe report count: {safeReportCount}")
