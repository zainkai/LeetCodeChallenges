function swap(arr, x, y) {
  let temp = arr[x]
  arr[x] = arr[y]
  arr[y] = temp
}


var sortColors = function(arr) {
  for(let i = arr.length -1; i >= 0 ; i--) {
    for(let j = i;j >= 0; j--) {
      if(i==j) continue
      if(arr[i] < arr[j]) {
         swap(arr, j, i)
      }
    }
  }
}