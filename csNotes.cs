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
