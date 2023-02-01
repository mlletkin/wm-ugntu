import math


def y(x):
    return (abs(math.sin(x**2+1)))**(1/3) - x


if __name__ == "__main__":
    a, b = 0, 1
    e = 0.000001
    c = (a+b)/2
    while True:
        if abs(b - a) < e:
            print(c)
            break
        print(y(a)*y(c))
        if y(a)*y(c) < 0:
            b = c
        else:
            a = c
        
        c = (a + b) / 2

     
    
