using System.Collections;
using System.Linq;

public class Solution {  
    public bool checkMatch(char x, char y) {
        bool valid = false;
        switch(x) {
            case '(':
                valid = y == ')';
                break;
            case '[':
                valid = y == ']';
                break;
            case '{':
                valid = y == '}';
                break;
            case ')':
                valid = y == '(';
                break;
            case ']':
                valid = y == '[';
                break;
            case '}':
                valid = y == '{';
                break;
            default:
                valid = false;
                break;
        }
        return valid;
    }
    
    public bool IsValid(string s) {
        var charStk = new Stack<char>();
        
        if(s.Length % 2 != 0) {
            return false;
        }
        
        foreach(var c in s) {
            if(charStk.Count() > 0 && checkMatch(c,charStk.Peek())){
                charStk.Pop();
            }
            else {
                charStk.Push(c);
            }
            
        }
        
        return charStk.Count() == 0;
    }
}
