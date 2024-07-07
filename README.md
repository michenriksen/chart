# chart

A command-line tool for rendering bar charts that can be displayed directly in the terminal or in text-based files like
Markdown.

## Usage

By default, `chart` accepts line-based data with a numeric value and label text separated by a space or common tabular
symbols. It also tolerates additional whitespace, empty lines, common punctuation, and currency symbols:

```console
$ cat examples/quarterly-stats.csv
18.67,2021-Q1
19.45,2021-Q2
20.89,2021-Q3
21.34,2021-Q4
21.78,2022-Q1
22.45,2022-Q2
23.89,2022-Q3
24.34,2022-Q4
23.54,2023-Q1
24.78,2023-Q2
25.89,2023-Q3
26.34,2023-Q4
28.15,2024-Q1
25.44,2023-Q2
$ chart -i examples/quarterly-stats.csv
2021-Q1 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 18.67
2021-Q2 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 19.45
2021-Q3 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 20.89
2021-Q4 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 21.34
2022-Q1 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 21.78
2022-Q2 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 22.45
2022-Q3 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 23.89
2022-Q4 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 24.34
2023-Q1 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 23.54
2023-Q2 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 25.44
2023-Q3 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 25.89
2023-Q4 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 26.34
2024-Q1 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 28.15
```

You can also chart the occurrences of lines using the `-c/--count` flag:

```console
$ cat examples/count-occurrences.txt
Three
Two
Four
Five
Four
One
Three
Three
Five
Two
Five
Four
Five
Four
Five
$ chart -i examples/count-occurrences.txt --count
Three ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 3
  Two ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 2
 Four ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 4
 Five ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 5
  One ▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 1
```

### Piping data

Pipes allow you to feed data to `chart`, making it easy to chart the output of other terminal commands. For example,
if you have a file with newline-delimited JSON data from a source code security scanner, you can chart statistics on
identified CWEs. Process the data with `jq`, count the lines with `sort` and `uniq`, and pipe it into `chart` for
visualization:

```console
$ jq -r '.cwe' examples/sast-findings.jsonld | sort | uniq -c | chart
CWE-200 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 2
 CWE-22 ▇▇▇▇▇▇▇▇▇▇ 1
 CWE-23 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-248 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-284 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-307 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-312 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 3
CWE-327 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 3
CWE-362 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-400 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-434 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-502 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-532 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 2
CWE-601 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-611 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-676 ▇▇▇▇▇▇▇▇▇▇ 1
 CWE-78 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 6
 CWE-79 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 7
CWE-798 ▇▇▇▇▇▇▇▇▇▇ 1
 CWE-89 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 7
CWE-918 ▇▇▇▇▇▇▇▇▇▇ 1
 CWE-94 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 5
```

### Sorting and ordering

By default, the chart displays bars in the order of insertion. This is usually fine, but sorting the chart differently
can make the data easier to read and understand. Let's order the CWE chart by label:

```console
$ jq -r '.cwe' examples/sast-findings.jsonld | sort | uniq -c | chart --sort label
CWE-200 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 2
 CWE-22 ▇▇▇▇▇▇▇▇▇▇ 1
 CWE-23 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-248 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-284 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-307 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-312 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 3
CWE-327 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 3
CWE-362 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-400 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-434 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-502 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-532 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 2
CWE-601 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-611 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-676 ▇▇▇▇▇▇▇▇▇▇ 1
 CWE-78 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 6
 CWE-79 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 7
CWE-798 ▇▇▇▇▇▇▇▇▇▇ 1
 CWE-89 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 7
CWE-918 ▇▇▇▇▇▇▇▇▇▇ 1
 CWE-94 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 5
```

The chart is now sorted alphabetically, but it's more natural for this data to be sorted by the numbers in the CWE
identifiers. By using `labelnum` as the sorting option, numeric characters are extracted from the labels and sorted as
integers:

```console
$ jq -r '.cwe' examples/sast-findings.jsonld | sort | uniq -c | chart --sort labelnum
 CWE-22 ▇▇▇▇▇▇▇▇▇▇ 1
 CWE-23 ▇▇▇▇▇▇▇▇▇▇ 1
 CWE-78 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 6
 CWE-79 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 7
 CWE-89 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 7
 CWE-94 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 5
CWE-200 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 2
CWE-248 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-284 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-307 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-312 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 3
CWE-327 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 3
CWE-362 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-400 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-434 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-502 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-532 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 2
CWE-601 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-611 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-676 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-798 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-918 ▇▇▇▇▇▇▇▇▇▇ 1
```

You can also sort the chart by value and specify the order. Let's sort by value in descending order to easily identify
the most and least common weaknesses:

```console
$ jq -r '.cwe' examples/sast-findings.jsonld | sort | uniq -c | chart --sort value --desc
 CWE-89 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 7
 CWE-79 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 7
 CWE-78 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 6
 CWE-94 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 5
CWE-327 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 3
CWE-312 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 3
CWE-532 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 2
CWE-200 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 2
CWE-918 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-798 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-676 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-611 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-601 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-502 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-434 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-400 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-362 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-307 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-284 ▇▇▇▇▇▇▇▇▇▇ 1
CWE-248 ▇▇▇▇▇▇▇▇▇▇ 1
 CWE-23 ▇▇▇▇▇▇▇▇▇▇ 1
 CWE-22 ▇▇▇▇▇▇▇▇▇▇ 1
```

### Scaling

Sometimes, smaller values are overshadowed by larger values in the distribution:

```console
$ cat examples/statistics.txt
1.2    Sales leads
3.8    Conversion rate
12.6   Monthly subs
34.9   Quarterly revenue
79.4   Client meetings
158.2  Product launches
367.1  Employees
952.5  Support calls
2764.8 Website traffic
7360.4 Marketing budget
$ chart -i examples/statistics.txt
      Sales leads ▏ 1.2
  Conversion rate ▏ 3.8
     Monthly subs ▏ 12.6
Quarterly revenue ▏ 34.9
  Client meetings ▇ 79.4
 Product launches ▇ 158.2
        Employees ▇▇▇ 367.1
    Support calls ▇▇▇▇▇▇▇ 952.5
  Website traffic ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 2764.8
 Marketing budget ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 7360.4
```

Using the `--scale` flag scales the chart logarithmically, improving readability in these situations:

```console
$ chart -i examples/statistics.txt --scale
      Sales leads ▇▇▇▇▇ 1.2
  Conversion rate ▇▇▇▇▇▇▇▇▇▇ 3.8
     Monthly subs ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 12.6
Quarterly revenue ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 34.9
  Client meetings ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 79.4
 Product launches ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 158.2
        Employees ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 367.1
    Support calls ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 952.5
  Website traffic ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 2764.8
 Marketing budget ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 7360.4
```

### Output formats

In addition to text-based charts, `chart` can generate Mermaid XY charts using the `--mermaid` flag:

```console
$ chart -i examples/count-occurrences.txt --count --sort value --mermaid --title 'Occurrences'
xychart-beta
  title "Occurrences"
  x-axis ["One", "Two", "Three", "Four", "Five"]
  bar [1, 2, 3, 4, 5]
```

See the [rendered chart here](https://mermaid.live/edit#pako:eNo9j70OwjAMhF_F8hwG_pbOiA0xwATpYBJDI2hSuQm0Qrw7KRS2891n2fdEEyxjgV1vKpI4OXEk7QGiizcGjVtjkgh7w63GIegm1LkWjjnyrFFlZv8Io6iER28dkozK3bNXDrsnEjhOFcwUzBUsFCxLVFiz1ORsfuI5QBpjxXVeKbK0JNfh7itzlGLY9d5gcaZbywolpEv1n1JjKfLK0UWo_rtsXQyy-Zb8dFXYkD-E8GNebwNZU54).

It can also generate a basic configuration for a Chart.js bar chart using the `--chartjs` flag:

```console
$ chart -i examples/count-occurrences.txt --count --sort value --chartjs --title 'Occurrences'
// Chart.js configuration (https://www.chartjs.org/docs/latest/configuration/).
// Generated by chart (https://github.com/michenriksen/chart).
const config = {
  type: "bar",
  data: {
    datasets: [{
      data: [1,2,3,4,5],
    }],
    labels: ["One","Two","Three","Four","Five"],
  },
  options: {
    plugins: {
      title: {
        display: true,
        text: "Occurrences"
      }
    }
  }
}
```

### Additional options

See `chart --help` for additional flags and options.

## Installation

Download the latest pre-compiled binary for your operating system from the [releases page].

If you have Go installed, you can also install the latest binary with:

```bash
go install github.com/michenriksen/chart/cmd/chart@latest
```

## Similar tools

- [termgraph]: A more powerful CLI tool, and the inspiration for `chart`, supporting multiple data sets, colors, and
heat maps.
- [spark]: Another CLI tool for rendering simple sparklines in the terminal.

[CWEs]: https://cwe.mitre.org/
[jq]: https://jqlang.github.io/jq/
[sort]: https://man7.org/linux/man-pages/man1/sort.1.html
[uniq]: https://www.man7.org/linux/man-pages/man1/uniq.1.html
[Mermaid]: https://mermaid.live/
[Chart.js]: https://www.chartjs.org/
[releases page]: https://github.com/michenriksen/chart/releases
[termgraph]: https://github.com/mkaz/termgraph
[spark]: https://github.com/holman/spark

