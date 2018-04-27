
function fact(n)
    if n == 0 then
        return 1
    else
        return n * fact(n-1)
    end
end

print("enter a number:")
local num = io.read("*number")
print(num)
if num then
    print(fact(num))
else
    print("xxx")
end
