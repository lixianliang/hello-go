
polyline = {
    color = "blue",
    thickness = 2,
    npoints = 4,
    {x = 0, y = 0},
    {x = -10, y = 0},
    {x = -10, y = 1},
    {x = 0, y = 1}
}

local polyline2 = {
    color = "blue",
    thickness = 2,
    npoints = 4
}

print(polyline[2].x)
print(type(polyline2))
print(polyline2)

