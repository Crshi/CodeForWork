
public class Solution {
    public int MinMutation(string start, string end, string[] bank) {
        var changeCount = 0;

        for(var i = 0;i < start.Length;i++ )
        {
            if(start[i] != end[i])
            {
                start[i] = end[i];

                if(bank.Contains(start))
                {
                    changeCount++;
                }
                else
                {
                    return -1;
                }
            }
        }
    }
}

public main(){
    var s = Solution();
    s.MinMutation("AACCGGTT","AACCGGTA",{"AACCGGTA"});
}