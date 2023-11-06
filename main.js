const crypto = require('crypto')
const algorithm = 'aes-256-cbc'

const main = () => {
    const key = '0123456789abcdef0123456789abcdef'
    const data = 'hello'

    //const encrypted = 'P3ZtfTkvWHxQc2VEeTEzcdnV/fS8QYcEt89xarDI4yo=';
    const encrypted = encrypt(data, key)
    const decrypted = decrypt(encrypted, key)
    console.log(decrypted)
}

const encrypt = (data, key) => {
    //const initVectorDec = generateIV()
    //const initVector = Buffer.from(initVectorDec, 'ascii').toString('utf-8')
    //console.log('iv : ', initVector)

    const initVector = 'F?H0GtWx?lFWnF1Q'

    const cipher = crypto.createCipheriv(algorithm, key, initVector)
    let encrypted = cipher.update(data, 'utf-8', 'binary')
    encrypted += cipher.final('binary')

    //const finalEncrypted = Buffer.concat([Buffer.from(initVectorDec, 'ascii'), Buffer.from(encrypted, 'ascii')])
    //return finalEncrypted.toString('base64')

    return Buffer.from(encrypted, 'binary').toString('base64')
    //return encrypted
}

const decrypt = (encrypted, key) => {
    const raw = base64ToArrayBuffer(encrypted)
    console.log(raw)

    const d = Buffer.from(encrypted, 'base64').toString('binary')
    console.log(d)

    const iv = 'F?H0GtWx?lFWnF1Q'

    //const ivRaw = raw.slice(0, 16)
    //const dataRaw = raw.slice(16)

    //const iv = Buffer.from(ivRaw, 'ascii').toString()
    //const data = Buffer.from(dataRaw, 'ascii').toString('base64')
    //const data = Buffer.from(raw, 'ascii').toString('base64')
    //const data = Buffer.from('base64').toString('hex')
    //const data = encrypted

    const data = Buffer.of(raw, 'binary').toString('base64')
    
    var decipher = crypto.createDecipheriv(algorithm, key, iv);
    var decrypted = decipher.update(data, 'binary', 'utf8');
    //var decrypted = decipher.update(encrypted, 'base64', 'utf8');
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
