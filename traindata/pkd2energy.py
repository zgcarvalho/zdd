from math import log
f = open("./pkd")
lines = f.readlines()
for l in lines:
    data = l.split()
    pkd = float(data[1])
    kd = 10**-pkd
    energy = 8.314*310.2*log(kd)
    print "{}\t{}".format(data[0],energy)
