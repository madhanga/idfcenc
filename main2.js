const crypto = require('crypto')
const algorithm = 'aes-128-cbc'

const main = () => {
    const key = '0123456789abcdef'
    const data = 'hello how?'

    //const encrypted = 'amtsQENRXV1tQ2xcMnkwNflBQrE+xLPBueNgWXRCfpw=';
    const encrypted = encrypt(data, key)
    console.log(encrypted)

    const decrypted = decrypt(encrypted, key)
    console.log(decrypted)
}

const encrypt = (data, key) => {
    const initVectorDec = 'qcBm;QH4X4`x:Y;_'//generateIV()
    const initVector = Buffer.from(initVectorDec, 'ascii').toString('utf-8')
    //console.log(initVector)

    const cipher = crypto.createCipheriv(algorithm, key, initVector)
    //cipher.setAutoPadding(true)
    let encrypted = cipher.update(data, 'utf-8', 'base64')
    encrypted += cipher.final('base64')

    return encrypted

}

const decrypt = (encrypted, key) => {
    const initVectorDec = 'qcBm;QH4X4`x:Y;_'//generateIV()
    var decipher = crypto.createDecipheriv(algorithm, key, initVectorDec);
    //decipher.setAutoPadding(false)
    var decrypted = decipher.update(encrypted, 'base64', 'utf8');
    decrypted += decipher.final('utf8');
    
    return decrypted
}

function generateIV() {
    const length = 16
    const lowAscii = 47 
    const highAscii = 126

    var ivBuffer = []
    for (let i = 0; i < length; i++) {
        const digit = lowAscii + Math.floor(Math.random() * (highAscii - lowAscii))
        ivBuffer.push(digit)
    }

    return ivBuffer
}

function base64ToArrayBuffer(base64) {
    var text = atob(base64);
    var decimalArray = []
    for (var i = 0; i < text.length; i++) {
        decimalArray.push(text.charAt(i).charCodeAt(0));
    }

    return decimalArray
}

main()