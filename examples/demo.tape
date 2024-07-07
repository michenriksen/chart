Require chart
Require jq
Require sort
Require uniq

Output demo.mp4
Output demo.gif

Set FontSize 16

Type "jq -r '.cwe' sast-findings.jsonld | sort | uniq -c | chart --sort labelnum" Enter
Sleep 2s

Hide
Type "clear" Enter
Show

Type "jq -r '.cwe' sast-findings.jsonld | sort | uniq -c | chart --sort value --desc" Enter
Sleep 1s
