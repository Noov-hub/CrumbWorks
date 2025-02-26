#nyoba list
mainlist = "dapak"
listdata = [10, 'umur', 'intel']
print(mainlist, 'adalah', listdata[2])
#nyoba del
del listdata[2]
print(listdata)

#nyoba append, pop dan insert
listdata.append('bgst')
print(listdata)
intel_adalah = listdata.pop()
print(listdata)
print('intel adalah orang yg', intel_adalah)
listdata.insert(-2, 'bukan_umur')
print(listdata)

#nyoba sort
list1 = ['asa', 'de', 'bambang']
list1.sort()
print(list1)
list1.clear()
print(list1)
list1.insert(0, mainlist)

#nyoba boolean
print(list1)
dapakintel = True
dapakpropleyer = False
print(mainlist, 'adalah intel, statemen = ', dapakintel)
print('apakah dapak pro pleyer, statemen=', dapakpropleyer)

#ngubah data melalui index tanpa insert
listdata[0] = 11
print(listdata)
listdata[2] = 13
print(listdata)

#nyoba tuple persamaan
dapak = ('bgst', 'intel', 'noob')
sifat, pekerjaan, skill = dapak
print('dapak')
print('sifat:', sifat)
print('pekerjaan:', pekerjaan)
print('skill:', skill)

