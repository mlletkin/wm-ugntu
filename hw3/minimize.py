import numpy as np
from scipy.optimize import fsolve

def system(x):
    return [np.sin(x[0]+2)-x[1]-2.5, x[0] + np.cos(x[1] + 2) - 0.5]

x0 = [0,0]
sol = fsolve(system, x0, xtol=1e-5)
print(sol)