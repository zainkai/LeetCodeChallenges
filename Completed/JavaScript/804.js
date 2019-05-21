/**
 * @param {string[]} words
 * @return {number}
 */
const charMap = {
  'a':".-",
  'b':"-...",
  'c':"-.-.",
  'd':"-..",
  'e':".",
  'f':"..-.",
  'g':"--.",
  'h':"....",
  'i':"..",
  'j':".---",
  'k':"-.-",
  'l':".-..",
  'm':"--",
  'n':"-.",
  'o':"---",
  'p':".--.",
  'q':"--.-",
  'r':".-.",
  's':"...",
  't':"-",
  'u':"..-",
  'v':"...-",
  'w':".--",
  'x':"-..-",
  'y':"-.--",
  'z':"--.."
}
function morsify(word) {
  if (word == null) {
    return ''
  }
  return word.toLowerCase().split('').map(char => charMap[char]).join('')
}
var uniqueMorseRepresentations = function(words) {
  if (!words || words.length === 0) return 0
  
  const morseMap = {}
  for(let word of words) {
    const morsedWord = morsify(word)
    morseMap[morsedWord] = (morseMap[morsedWord] || 0) + 1
  }
  return Object.keys(morseMap).length
};