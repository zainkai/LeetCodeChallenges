# Write your MySQL query statement below
WITH emp_manger_salarys_tbl AS (
  SELECT
  e1.Name as Name,
  e1.Salary as emp_salary,
  e2.Salary as mgr_salary
  FROM Employee as e1
  JOIN Employee as e2 ON e1.ManagerId=e2.Id
)
SELECT
Name as Employee
FROM emp_manger_salarys_tbl
WHERE emp_salary > mgr_salary
