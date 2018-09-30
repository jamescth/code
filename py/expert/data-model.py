#!/usr/bin/env python3
# google search python data model for methods
# some behaviour that I want to implement -> write some __function__
# top-level function or top-level syntax -> corresponding __
#   x + y -> __add__
#   intit x -> __init
#   repr(x) -> __repr__
#   x()     -> __call__
#
# python3 -i data-model.py
# >>> p1
# >>> p2
# 
class Polynomial:
    def __init__(self, *coeffs):
        self.coeffs = coeffs

    def __repr__(self):
        return 'Polynomial(*{!r})'.format(self.coeffs)

    def __add__(self, other):
        return Polynomial(*(x+y for x, y in zip(self.coeffs, other.coeffs)))

    def __len__(self):
        return len(self.coeffs)

    def __call__(self):
        pass

p1 = Polynomial(1,2,3)
p2 = Polynomial(3,4,3)

