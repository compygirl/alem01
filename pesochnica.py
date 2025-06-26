import sys
import random

class City:
    def __init__(self, line):
        self.owner = line.split()[0]
        self.pop = int(line.split()[1])
        self.name = line.split()[2]
        self.x = int(line.split()[3])
        self.y =int (line.split()[4])
        self.a1 = int(line.split()[5])
        self.a2 = int(line.split()[6])
        
    def __str__(self):
        return "owner:"+self.owner+"; pop:"+str(self.pop)+"; name:"+self.name

    def attack(self, city):
        print(str(self.x)+" "+str(self.y)+" "+str(city.x)+" "+str(city.y), flush=True)
        print(str(self.x)+" "+str(self.y)+" "+str(city.x)+" "+str(city.y), file=sys.stderr, flush=True)


class World:
    def __init__(self, line):
        print(line, file=sys.stderr, flush=True)
        self.num_city = int(line.split()[0])
        self.my_name = line.split()[1]
        self.tick = int(line.split()[2])


while True:
    city_name = {}
    city_owner = {}
    is_printed = False
    
    wrld = World(input())

    for i in range(wrld.num_city):
        c = City(input())
        city_name[c.name] = c
        if not (c.owner in city_owner.keys()):
            city_owner[c.owner] = []
        city_owner[c.owner].append(c)

    m = int(input())

    for i in range(m):
        line = input()
        
    for c in city_owner[wrld.my_name]:
        print(c, file=sys.stderr, flush=True)
        if c.pop > 10:
            c.attack(city_owner["Neutral"][0])
            is_printed=True
            
    if not is_printed:
        print("", flush=True)
    """
    # use file=sys.stderr to print for debugging
    print("debug code", file=sys.stderr, flush=True)

    # this will choose one of random actions
    print("100 100 200 200", flush=True)
    """