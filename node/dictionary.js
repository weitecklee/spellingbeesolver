const fs = require("fs");
const path = require("path");

const wordList = fs
  .readFileSync(path.join(__dirname, "..", "wordlist.txt"), "utf-8")
  .split("\n");

class Trie {
  constructor() {
    this.isWord = false;
    this.children = new Map();
  }

  addWord(word) {
    let curr = this;
    for (const c of word) {
      if (!curr.children.has(c)) curr.children.set(c, new Trie());
      curr = curr.children.get(c);
    }
    curr.isWord = true;
  }

  findWords(letters, words, triePath, centerLetter, containsCenterLetter) {
    if (this.isWord && containsCenterLetter) words.push(triePath.join(""));
    for (const c of letters) {
      if (this.children.has(c)) {
        if (c === centerLetter) {
          this.children
            .get(c)
            .findWords(letters, words, triePath.concat(c), centerLetter, true);
        } else {
          this.children
            .get(c)
            .findWords(
              letters,
              words,
              triePath.concat(c),
              centerLetter,
              containsCenterLetter
            );
        }
      }
    }
  }
}

class Dictionary {
  constructor() {
    this.root = new Trie();
  }

  addWord(word) {
    this.root.addWord(word);
  }

  findWords(letters) {
    const words = [];
    const centerLetter = letters[0];
    this.root.findWords(letters, words, [], centerLetter, false);
    const pangramRegex = new RegExp(
      "^" +
        letters
          .split("")
          .map((c) => `(?=.*${c})`)
          .join("")
    );
    const pangrams = [];
    const others = [];
    for (const word of words) {
      if (pangramRegex.test(word)) pangrams.push(word.toUpperCase());
      else others.push(word);
    }
    pangrams.sort();
    others.sort();
    return pangrams.concat(others);
  }
}

const dictionary = new Dictionary();
for (const word of wordList) {
  if (word.length >= 4) dictionary.addWord(word);
}

module.exports = dictionary;
