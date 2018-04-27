
local ffi = require("ffi")

ffi.cdef[[
struct user {
    int age;
};

void foo_a(int len);
void foo_b(char *buf);
void foo_c(int *len, char *buf);
void foo_d(struct user *u);
]]


local foo = ffi.load("libfoo")
local length = ffi.new("int", 200)
str = "haha, this is a 12345"
local buf = ffi.new('unsigned char[?]', string.len(str), str)
local len = ffi.new("int[1]", string.len(str))

local user = ffi.new("struct user")
user.age = 8765

foo.foo_a(length)
foo.foo_b(buf)
foo.foo_c(len, buf)
print("foo_c", tonumber(len[0]))
print(ffi.string(buf))
foo.foo_d(user)
print("foo_d", user.age)



