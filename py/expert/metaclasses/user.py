#!/usr/bin/env python3
# user.py
from library import Base

# this is to verify if the Base contains foo before runtime
assert hasattr(Base, 'foo'), "you broke it, you fool!"

class Derived(Base):
    def bar(self):
        return 'bar'
