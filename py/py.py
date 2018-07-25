#!/usr/bin/env python
res =[]
print(res)
for i in range(3):
    res.append(i)
print("res range 3", res)

print("range(1,6,2)")
for i in range(1, 6, 2):
    print(i)

my_list = []
my_list.append("1")
my_list.insert(0, "2")
my_list.insert(0, "3")
# del my_list[-1]
# my_list.remove("2")
print("for l in my_list sort", sorted(my_list))

# for l in my_list:
#    print(l)

squares = [x**2 for x in range (1, 8)]
print("list comprehansion", squares)

l = [1,2,3]
if not l:
    print("l is NoneType")
else:
    print("l is", type(l))

# dictionary
alien = {'color':'green','points':5}
alien['new'] = 'add new k/v'
alien[0] = 0

for k, v in alien.items():
    print("%s" % k, v)

choices = ['red', 'green', 'blue']
i = 0
while i < len(choices):
    print('{}) {}'.format(i, choices[i]))
    i += 1

choices = ['red', 'green', 'blue']
for idx, choice in enumerate(choices):
    print('{}) {}'.format(idx, choice))
