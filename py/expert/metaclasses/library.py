#!/usr/bin/env python3
# library.py

# read doc derived type
# $ python3 -i user.py 
# BaseMeta.__new__ <class 'library.BaseMeta'> Base () {'__module__': 'library', '__qualname__': 'Base', 'foo': <function Base.foo at 0x7f38104fa840>}
# BaseMeta.__new__ <class 'library.BaseMeta'> Derived (<class 'library.Base'>,) {'__module__': '__main__', '__qualname__': 'Derived', 'bar': <function Derived.bar at 0x7f38104fa510>}
class BaseMeta(type):
    def __new__(cls, name, bases, body):
        if name != 'Base' and not 'bar' in body:
            # check inheritance obj has bar method
            raise TypeError("bad user class")
        print('BaseMeta.__new__', cls, name, bases, body)
        return super().__new__(cls, name, bases, body)

class Base(metaclass=BaseMeta):
    def foo(self):
        return self.bar

'''
class Base:
    def foo(self):
        return self.bar

# what the following code does:
#   $ python3 -i user.py
#   my buildclass-> <function Derived at 0x7fbb29274e18> Derived <class 'library.Base'> {}
#   check if bar method defined
#   my buildclass-> <function Completer at 0x7fbb2747f950> Completer None {}
old_bc = __build_class__
def my_bc(fun, name, base=None, **kw):
    print('my buildclass->', fun, name, base, kw)
    if base is Base:
        print('check if bar method defined')
    if base is not None:
        return old_bc(fun, name, base, **kw)
    return old_bc(fun, name, **kw)
import builtins
builtins.__build_class__ = my_bc
'''
