#!/usr/bin/env python3
"""
>>> add
<function add at 0x7ff17c9b6e18>
>>> add.__name__
'add'
>>> add.__module__
'__main__'
>>> add.__defaults__
(10,)
>>> add.__defaults__
(10,)
>>> add.__code__
<code object add at 0x7f7fce2be810, file "dec.py", line 3>
>>> add.__code__.co_code
b'|\x00|\x01\x17\x00S\x00'
>>> add.__code__.co_nlocals
2
>>> add.__code__.co_varnames
('x', 'y')

>>> from inspect import getsource
>>> getsource(add)
'def add(x, y=10):\n    return x + y\n'
>>> print(getsource(add))
def add(x, y=10):
    return x + y

>>> from inspect import getfile
>>> getfile(add)
'dec.py'

add.__{tab tab}
"""
from time import time

def timer(func):
    def f(*args, **kwargs):
        before = time()
        rv = func(*args, **kwargs)
        after = time()
        print('elapsed', after-before)
        return rv
    return f

def ntimes(n):
    def inner(f):
        def wrapper(*args, **kwargs):
            for _, in range(n):
                print('running {.__name__}'.format(f))
                rv = f(*args, **kwargs)
            return rv
        return wrapper
    return inner

@timer
def add(x, y=10):
    return x + y

@timer
def sub(x, y=10):
    return x - y

print('add(10)',       add(10))
print('add(20, 30)',   add(20, 30))
print('sub(10)',       sub(10))
print('sub(20, 30)',   sub(20, 30))

