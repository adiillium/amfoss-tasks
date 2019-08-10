string = input()
string_length = len(string)
current_count = 1
for i in range(1, string_length):
    if string[i] == string[i - 1]:
        current_count += 1
    else:
        current_count = 1
    if current_count >= 7:
        print('YES')
        exit()
print('NO')
