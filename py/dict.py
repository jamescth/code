#!/usr/bin/env python

users = [
        {
            'last':'fermi',
            'first':'enrico',
            'username':'efermi',
        },
        {
            'last':'curie',
            'first':'marie',
            'username':'mcurie',
        },
]

for user in users:
    for k, v in user.items():
        print(k + ": " + v)
    print("")
