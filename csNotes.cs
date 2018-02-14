// var   for quicker programming.
// pibbXtra will alway keep type List<int>
var pibbXtra = new List<int>();

// Write to Console.
Console.WriteLine("INT|STRING|BOOL");

// NoCoalesce set to val if null
MyClass temp = x ?? new MyClass();

//helpful Lambdas

//  .Where(x => x) returns IEnumerable.
List<int> ids = new List(genIds());
List<int> lessThan100 = ids.Where(x => x < 100).ToList();
ids.add(23423);


// Type Conversions
// String to int
int x =0;
if (Int32.TryParse("323123", out x)) { }
int x = Int32.Parse("23123");

// int to string
int x = 3324234;
string str = x.ToString();

//String to char Pretty simple
string str = "blahsdadsa";
foreach(char i in str) {
    //do stuff
}
char x = str[2];

// ADT's
using System.Collections;

Stack<T> stk = new Stack<T>();
stk.Count;
stk.Clear();
stk.Contains('Blah');
stk.Peek(); // like top
stk.Pop();
stk.Push('asdasd');
