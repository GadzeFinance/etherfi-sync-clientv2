const crypto = require('crypto')

// Encrypt any text (used to encrypt validator keys and passwords)
function encrypt(text, ENCRYPTION_KEY) {
  //const IV_LENGTH = 16; // For AES, this is always 16
  //let iv = crypto.randomBytes(IV_LENGTH);

  //console.log("iv: ", iv.toString('hex'))

  // iv test
  const iv = Buffer.from("4fa9077c8f4ae788988bb8cae303bf53", "hex")

  let cipher = crypto.createCipheriv('aes-256-cbc', Buffer.from(ENCRYPTION_KEY), iv);
  let encrypted = cipher.update(text);

  encrypted = Buffer.concat([encrypted, cipher.final()]);

  // console.log("cipher.final():", cipher.final().toString('hex'))

  return iv.toString('hex') + ':' + encrypted.toString('hex');
}

const key = "6368616e676520746869732070617373"
const text = "exampleplaintext"
console.log(encrypt(text, key))