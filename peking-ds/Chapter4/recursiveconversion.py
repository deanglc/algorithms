

def toStr(n, base):
    convertString = "0123456789ABCDEF"
    if n < base:
        return convertString[n]
    else:
        return toStr(n // base, base) + convertString[n % base]


print(toStr(1453, 16))
print(toStr(769, 10))
print(toStr(3, 2))
