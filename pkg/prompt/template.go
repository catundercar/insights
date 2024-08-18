package prompt

const promptTemplate = `Please analyze the following slow query log from {{.DatabaseName}} database and provide a detailed problem diagnosis and optimization suggestions. Use {{.Language}} and strictly follow this format for your analysis output:

## 1. Query Overview
[Briefly describe the purpose and main operations of the query]

## 2. Performance Issues
[List and explain the main problems causing the query to be slow, such as:
- Improper use of indexes
- Low efficiency of table joins
- Large data scan volume
- Poor subquery performance
- Other relevant issues]

## 3. Detailed Analysis
[Provide an in-depth analysis of each identified problem, including:
- Specific manifestation of the problem
- Possible causes
- Degree of impact on query performance]

## 4. Optimization Suggestions
[Offer specific optimization suggestions for each problem, such as:
- Index optimization strategies
- Query rewrite suggestions
- Table structure adjustments
- Configuration parameter tuning
- Other relevant recommendations]

## 5. Expected Results
[Describe the expected performance improvements after implementing the optimization suggestions]

## 6. Additional Recommendations
[Provide any other suggestions or observations that might help improve overall database performance]

Please ensure your analysis is comprehensive, accurate, and provides actionable specific recommendations.

Slow Query Log:
{{.SlowQueryLog}}`
