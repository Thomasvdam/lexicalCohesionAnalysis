#!/usr/bin/env python
import csv
import matplotlib.pyplot as plt
import sys

dataPath = sys.argv[1]

res = csv.reader(open(dataPath), delimiter=',')

values = []
for col in res:
  values.append(int(col[1]))

plt.rcParams['figure.figsize'] = len(values), 10

N = len(values)
x = range(N)
width = 1
plt.bar(x, values, width, color="blue")

plt.show()
