<p>给你一个二维整数数组&nbsp;<code>queries</code>&nbsp;，其中 <code>queries[i] = [n<sub>i</sub>, k<sub>i</sub>]</code> 。第&nbsp;<code>i</code>&nbsp;个查询&nbsp;<code>queries[i]</code> 要求构造长度为&nbsp;<code>n<sub>i</sub></code> 、每个元素都是正整数的数组，且满足所有元素的乘积为&nbsp;<code>k<sub>i</sub></code><sub>&nbsp;</sub>，请你找出有多少种可行的方案。由于答案可能会很大，方案数需要对 <code>10<sup>9</sup> + 7</code>&nbsp;<strong>取余</strong> 。</p>

<p>请你返回一个整数数组<em>&nbsp;</em><code>answer</code>，满足<em>&nbsp;</em><code>answer.length == queries.length</code>&nbsp;，其中<em>&nbsp;</em><code>answer[i]</code>是第<em>&nbsp;</em><code>i</code>&nbsp;个查询的结果。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<b>输入：</b>queries = [[2,6],[5,1],[73,660]]
<b>输出：</b>[4,1,50734910]
<b>解释：</b>每个查询之间彼此独立。
[2,6]：总共有 4 种方案得到长度为 2 且乘积为 6 的数组：[1,6]，[2,3]，[3,2]，[6,1]。
[5,1]：总共有 1 种方案得到长度为 5 且乘积为 1 的数组：[1,1,1,1,1]。
[73,660]：总共有 1050734917 种方案得到长度为 73 且乘积为 660 的数组。1050734917 对 10<sup>9</sup> + 7 取余得到 50734910 。
</pre>

<p><strong>示例 2&nbsp;：</strong></p>

<pre>
<b>输入：</b>queries = [[1,1],[2,2],[3,3],[4,4],[5,5]]
<b>输出：</b>[1,2,3,10,5]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= queries.length &lt;= 10<sup>4</sup> </code></li> 
 <li><code>1 &lt;= n<sub>i</sub>, k<sub>i</sub> &lt;= 10<sup>4</sup></code></li> 
</ul>

<div><li>👍 35</li><li>👎 0</li></div>