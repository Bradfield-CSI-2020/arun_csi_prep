const faker = require('faker')

const primesInOrder = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199];
const alphabets = "abcdefghijklmnopqrstuvwxyz". split("");


const primeValueMap = new Map()

alphabets.forEach((letter, index) => {
    primeValueMap.set(letter, primesInOrder[index])
})

function anagramCheck(string1, string2) {
    
    if (string1.length !== string2.length) {
        return false
    }

    let string1Value = 1
    let string2Value = 1

    for(let i = 0; i < string1.length;  i++) {
        string1Value = string1Value * primeValueMap.get(string1[i])
        string2Value = string2Value * primeValueMap.get(string2[i])
    }

    console.log(string1Value)
    console.log(string2Value)
    return string1Value === string2Value

}

console.log(anagramCheck("hello", "helso"))
console.log(anagramCheck("hello", "hello"))


const word = 'radiometeorograph'.split('').sort().join('')
console.log(anagramCheck('radiometeorograph', word)) // doesn't work because of js precision issues

// todo: understand the js precision issue better
// todo: write a safe multiply function? unrelated to this problem