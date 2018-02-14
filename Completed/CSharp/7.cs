using System.Linq;

public class Solution {
    public int Reverse(int x) {
        string y;
        int newNum;
        int sign = 1;
        
        //Edge case
        if(x > Int32.MaxValue || x < Int32.MinValue) {
            return 0;
        }
        
        //Neg case
        if(x < 0) {
            sign = -1;
            x *=-1;
        }
        
        //ugly reversal
        y = new string(x.ToString().ToCharArray().Reverse().ToArray());
        
        //Parsing
        if(Int32.TryParse(y, out newNum)) {
            return newNum *sign;
        } else {
            return 0;
        }
    }
}
