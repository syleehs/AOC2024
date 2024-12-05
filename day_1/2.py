def pt2():
    d = {}
    res = 0
    with open("input", "r") as file:
        for line in file:
            a=line.split()
            if a[1] in d:
                d[a[1]]+=1
            else:
                d[a[1]]=1
    with open("input", "r") as file:
        for line in file:
            a=line.split()
            if a[0] in d:
                res+=int(a[0])*d[a[0]]
    return res
    
def main():
    print(pt2())

if __name__ == '__main__':
    main()