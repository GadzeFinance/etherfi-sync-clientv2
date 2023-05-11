const crypto = require('crypto')

function decrypt(text, ENCRYPTION_KEY) {

  let textParts = text.split(':');
  let iv = Buffer.from(textParts.shift(), 'hex');
  let encryptedText = Buffer.from(textParts.join(':'), 'hex');
  let decipher = crypto.createDecipheriv('aes-256-cbc', Buffer.from(ENCRYPTION_KEY), iv);
  let decrypted = decipher.update(encryptedText);

  decrypted = Buffer.concat([decrypted, decipher.final()]);

  return decrypted.toString();
}

const key = "7630bac8401d77beeb16f0c64815178bdb6f621ec694ab3dd8cfb3da8541a3e4"
const text = "94fa0973a57bfc09a676a4f2ab69b6d7:e8fa115d66beebf5e9df1cc781e357f6325c12737a8035444498e89e2dac9af4"
console.log(Buffer.from(key, 'hex'))
console.log(decrypt(text, key))