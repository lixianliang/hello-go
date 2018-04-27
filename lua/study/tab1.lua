days = {"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
print(days[1]) --> sunday

tab = {math.sin(1), math.sin(2)}
print(tab)

a = {x=0, y=0} 
print("xxxx")
print(a)
print(unpack(a))

w = {x=0, y=0, label="console"}
x = {math.sin(0), math.sin(1), math.sin(2)}
w[1] = "another field"
print(w[1][2])
x.f = w
print(w["x"])
print(w[1])
w.x = nil
