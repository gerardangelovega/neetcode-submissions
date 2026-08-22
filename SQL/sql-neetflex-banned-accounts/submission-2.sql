SELECT DISTINCT l1.account_id
FROM log_info AS l1
JOIN log_info AS l2 ON l1.account_id = l2.account_id
WHERE l1.ip_address != l2.ip_address
  AND l1.login <= l2.logout
  AND l2.login <= l1.logout
;