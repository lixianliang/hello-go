
function list_iter(t)
    local i = 0
    local n = table.getn(t)
    return function()
        i  = i + 1
        if i <= n then return t[i] end
    end
end

t = {10, 20, 30}
iter = list_iter(t)
while true do
    local element = iter()
    if element == nil then break end
    print(element)
end

t = {30, 40, 50}
for element in list_iter(t) do 
    print(element)
end
