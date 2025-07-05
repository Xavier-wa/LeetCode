import sys
s = input()

cnt_0 = 0 
cnt_1 = 0
ans = 1
for i,c in enumerate(s):
    if c == '0':
        ans += cnt_1
        cnt_0 += 1
    else:
        ans += cnt_0
        cnt_1 += 1
print(ans)