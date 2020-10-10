SELECT Users.name as NAME, SUM(Transactions.amount) as BALANCE
FROM Users INNER JOIN Transactions ON Transactions.account=Users.account
GROUP BY Users.account
HAVING BALANCE > 10000
