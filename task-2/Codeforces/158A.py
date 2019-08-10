n, k = map(int, input().split())
a = list(map(int, input().split()))
positives = sum(ai > 0 for ai in a)
if positives < k:
    print(positives)
else:
    kths = a.count(a[k - 1])
    greaters = sum(ai > a[k - 1] for ai in a)
    print(kths + greaters)
