public class Solution {
	public int HammingWeight(uint n) {
		var count = 0;//统计1的个数
		while(n>0)
		{
			n = n & (n-1);
			count ++;
		}
		return count;
	}
}