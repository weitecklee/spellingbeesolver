const dictionary = require("./dictionary");

const letters = process.argv[2];

if (!letters) {
  console.error("No letters provided.");
  return;
}

for (const word of dictionary.findWords(letters)) {
  console.log(word);
}
