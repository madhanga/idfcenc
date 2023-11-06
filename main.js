const crypto = require('crypto')
const algorithm = 'aes-256-cbc'

const main = () => {
    const key = '0123456789abcdef0123456789abcdef'

    if (process.argv.length != 4) {
        console.log('node main.js <encrypt|decrypt> "<data>"')
        return
    }

    const kind = process.argv[2]
    if (kind === 'encrypt') {
        console.log(encrypt(process.argv[3], key))
        return
    }

    if (kind === 'decrypt') {
        console.log(decrypt(process.argv[3], key))
        return
    }

    console.log('node main.js <encrypt|decrypt> "<data>"')
}

const encrypt = (data, key) => {
    const initVectorDec = generateIV()
    const initVector = Buffer.from(initVectorDec, 'ascii').toString('utf-8')

    const cipher = crypto.createCipheriv(algorithm, key, initVector)
    let encrypted = cipher.update(data, 'utf-8', 'binary')
    encrypted += cipher.final('binary')

    const finalEncrypted = Buffer.concat([Buffer.from(initVectorDec, 'ascii'), Buffer.from(encrypted, 'binary')])
    return Buffer.from(finalEncrypted, 'binary').toString('base64')
}

const decrypt = (encrypted, key) => {
    const raw = Buffer.from(encrypted, 'base64').toString('binary')

    const ivRaw = raw.slice(0, 16)
    const dataRaw = raw.slice(16)
    const initVector = Buffer.from(ivRaw, 'binary').toString()
    const data = Buffer.from(dataRaw, 'binary').toString('base64')

    var decipher = crypto.createDecipheriv(algorithm, key, initVector);
    var decrypted = decipher.update(data, 'base64', 'utf8');
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

/*function base64ToArrayBuffer(base64) {
    var text = atob(base64);
    var decimalArray = []
    for (var i = 0; i < text.length; i++) {
        decimalArray.push(text.charAt(i).charCodeAt(0));
    }

    return decimalArray
}*/

main()
