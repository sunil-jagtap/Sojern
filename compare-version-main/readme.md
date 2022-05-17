## Scripting - 'compare versions'

Compare two version numbers version1 and version2.

- If version1 > version2 return 1
- If version1 < version2 return -1
- otherwise return 0


You may assume that the version strings are non-empty and contain only digits and the 'dot'
character. The 'dot' character does not represent a decimal point and is used to separate
number sequences. For instance '2.5' is not "two and a half" or "half way to version three", it is
the fifth second-level revision of the second first-level revision.


Here is an example of version numbers ordering: 0.1 < 1.1 < 1.2 < 1.2.9.9.9.9 < 1.3 < 1.3.4 <
1.10

For test,
    need to run main.go and hit http://localhost:8082/compare?s=1.2.4.0&c=1.2.3.4 on local