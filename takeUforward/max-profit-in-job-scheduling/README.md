Maximum Profit in Job Scheduling

You are given n jobs, where every job is represented by:

startTime[i]: the start time of the job,
endTime[i]: the end time of the job,
profit[i]: the profit earned from completing the job.


You cannot take two jobs that overlap in time.



Return the maximum profit you can earn such that there are no two overlapping jobs in your selected subset.



Note: A job ending at time X is allowed to overlap with another job that starts exactly at time X.


Examples:
Input: startTime = [1, 2, 3, 3], endTime  = [3, 4, 5, 6], profit  = [50, 10, 40, 70]

Output: 120

Explanation:

Select jobs [0, 3] with intervals [1-3] and [3-6].
Profit = 50 + 70 = 120.
Input: startTime = [1, 2, 3, 4, 6], endTime  = [3, 5, 10, 6, 9], profit  = [20, 20, 100, 70, 60]

Output: 150

Explanation:

Select jobs [0, 3, 4]: [1-3], [4-6], [6-9]
Total profit = 20 + 70 + 60 = 150.