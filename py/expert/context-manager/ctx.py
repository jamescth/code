#!/usr/bin/env python3
from sqlite3 import connect

'''
with ctx() as x:
    pass

x = ctx().__enter__
try:
    pass
finally:
    x.__exit__

'''


'''
class contextmanager:
    def __init__(self, cur):
        self.cur = cur
    def __enter__(self):
        print('__enter__')
        self.cur.execute('create table points(x int, y int)')
    def __exit__(self, *args):
        print('__exit__')
        self.cur.execute('drop table points')
'''

# this same as
# from contextlib import contextmanager
class contextmanager:
    def __init__(self, gen):
        self.gen = gen
    def __call__(self, *args, **kwargs):
        self.args, self.kwargs = args, kwargs
        return self
    def __enter__(self):
        self.gen_inst = self.gen(*self.args, **self.kwargs)
        next(self.gen_inst)
    def __exit__(self, *args):
        next(self.gen_inst, None)

@contextmanager
def temptable(cur):
    print("create table")
    cur.execute('create table points(x int, y int)')
    try:
        yield
    finally:
        print("drop table")
        cur.execute('drop table points')
#temptable = contextmanager(temptable)

with connect('test.db') as conn:
    cur = conn.cursor()
    #with contextmanager(temptable)(cur):
    with temptable(cur):
        cur.execute("insert into points (x, y) values(1, 1)")
        cur.execute("insert into points (x, y) values(1, 2)")
        cur.execute("insert into points (x, y) values(2, 1)")
        for row in cur.execute('select x, y from points'):
            print(row)
        for row in cur.execute('select sum(x * y) from points'):
            print(row)
