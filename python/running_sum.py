def running_sum(nums):
    running_total = []
    total = nums[0]
    running_total.append(total)

    for i in range(1, len(nums)):
        total += nums[i] 
        running_total.append(total)
    return running_total