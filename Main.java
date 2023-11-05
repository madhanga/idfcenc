import java.util.Arrays;
import java.util.Base64;
import java.util.HexFormat;
import java.util.Random;

import javax.crypto.Cipher;
import javax.crypto.spec.IvParameterSpec;
import javax.crypto.spec.SecretKeySpec;
 
public class Main {
    public static void main(String[] args) throws Exception {
        var encryptedFromGo = "aVRxW1pDL2V7en5idnZzZMsccCua4ZYQ4ZZpbWetlVGk9y0IahiwQJJJoCeve0u305sEgP9KE5rEEfwf5Upm1B6RSriDGQk7uSznOCYFkv3BYRCeGvZ2c1u+q5/mhTAw";
        if (args.length > 0) {
            encryptedFromGo = args[0];
        }

        final var keyStr = "0123456789abcdef";
        final var key = HexFormat.of().formatHex(keyStr.getBytes());
        final var decrypted = decrypt(encryptedFromGo, key);
        System.out.println(decrypted);
    }


    public static String generateIv() {
        final int lowAsciiLimit = 47; // '/'
        final int highAsciiLimit = 126; // 'z'
        
        final int ivLength = 16;
        final var random = new Random();
        final var ivBuffer = new StringBuilder(ivLength);
        
        for (int i=0; i<ivLength; i++) {
            final int randomNumber = lowAsciiLimit + (int) (random.nextFloat() * (highAsciiLimit - lowAsciiLimit));
            ivBuffer.append((char) randomNumber);
        }
        return ivBuffer.toString();
    }

    public static String encrypt(String data, String key) throws Exception {
        final var keyInHexBytes = HexFormat.of().parseHex(key);
        final var initVector = generateIv();
        final var ivSpec = new IvParameterSpec(initVector.getBytes("UTF-8"));
        final var keySpec = new SecretKeySpec(keyInHexBytes, "AES");
        final var ciper = Cipher.getInstance("AES/CBC/PKCS5PADDING");
        ciper.init(Cipher.ENCRYPT_MODE, keySpec, ivSpec);

        final var encryptedBytes = ciper.doFinal(data.toString().getBytes());
        final var finalArray = new byte[initVector.length() + encryptedBytes.length];
        
        System.arraycopy(initVector.getBytes(), 0, finalArray, 0, initVector.getBytes().length);
        System.arraycopy(encryptedBytes, 0, finalArray, initVector.getBytes().length, encryptedBytes.length);
        
        final var encrypted = Base64.getEncoder().encodeToString(finalArray);
        return encrypted;
    }

    public static String decrypt(String encrypted, String key) throws Exception {
        final var keyInHeBytes = HexFormat.of().parseHex(key);
        final var keySpec = new SecretKeySpec(keyInHeBytes, "AES");

        final var encryptedBytes = Base64.getDecoder().decode(encrypted);
        final var initVector = Arrays.copyOfRange(encryptedBytes, 0, 16);
        final var encryptedPayload = Arrays.copyOfRange(encryptedBytes, initVector.length, encryptedBytes.length);

        final var ivSpec = new IvParameterSpec(initVector);
        final var ciper = Cipher.getInstance("AES/CBC/PKCS5PADDING");
        ciper.init(Cipher.DECRYPT_MODE, keySpec, ivSpec);

        final var decryptedBytes = ciper.doFinal(encryptedPayload);
        final var decrypted = new String(decryptedBytes);

        return decrypted;
    }
}

